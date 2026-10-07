package cinema

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type transactionRunner interface {
	WithinTransaction(context.Context, func(context.Context, pgx.Tx) error) error
}

type screeningSeatTaker interface {
	TakeSeat(context.Context, DBTX, int64) error
}

type bookingCreator interface {
	Create(context.Context, DBTX, int64, int64) (Booking, error)
}

type BookingService struct {
	txManager           transactionRunner
	screeningRepository screeningSeatTaker
	bookingRepository   bookingCreator
}

func NewBookingService(
	txManager transactionRunner,
	screeningRepository screeningSeatTaker,
	bookingRepository bookingCreator,
) *BookingService {
	return &BookingService{
		txManager:           txManager,
		screeningRepository: screeningRepository,
		bookingRepository:   bookingRepository,
	}
}

func (s *BookingService) Book(
	ctx context.Context,
	screeningID int64,
	userID int64,
) (Booking, error) {
	var booking Booking
	err := s.txManager.WithinTransaction(ctx, func(txCtx context.Context, tx pgx.Tx) error {
		if err := s.screeningRepository.TakeSeat(txCtx, tx, screeningID); err != nil {
			return err
		}

		created, err := s.bookingRepository.Create(txCtx, tx, screeningID, userID)
		if err != nil {
			return err
		}
		booking = created
		return nil
	})
	if err != nil {
		return Booking{}, err
	}
	return booking, nil
}
