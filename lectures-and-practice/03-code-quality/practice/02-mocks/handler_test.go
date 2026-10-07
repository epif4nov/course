package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"code-quality-practice/02-mocks"
	"code-quality-practice/02-mocks/mocks"
	"go.uber.org/mock/gomock"
)

func TestHandler(t *testing.T) {
	t.Run("successful greeting", func(t *testing.T) {
		controller := gomock.NewController(t)
		service := mocks.NewMockGreeter(controller)
		service.EXPECT().
			Greet(gomock.Any(), "Ada").
			Return("Hello, Ada!", nil)

		request := httptest.NewRequest(http.MethodGet, "/greet?name=Ada", nil)
		recorder := httptest.NewRecorder()
		httpHandler := handler.NewHandler(service)

		httpHandler.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
		}
		if got, want := recorder.Body.String(), "Hello, Ada!"; got != want {
			t.Errorf("body = %q, want %q", got, want)
		}
	})

	t.Run("missing name", func(t *testing.T) {
		controller := gomock.NewController(t)
		service := mocks.NewMockGreeter(controller)

		request := httptest.NewRequest(http.MethodGet, "/greet", nil)
		recorder := httptest.NewRecorder()
		handler.NewHandler(service).ServeHTTP(recorder, request)

		if recorder.Code != http.StatusBadRequest {
			t.Fatalf("status = %d, want %d", recorder.Code, http.StatusBadRequest)
		}
	})

	t.Run("service error", func(t *testing.T) {
		controller := gomock.NewController(t)
		service := mocks.NewMockGreeter(controller)
		service.EXPECT().
			Greet(gomock.Any(), "blocked").
			Return("", errors.New("blocked"))

		request := httptest.NewRequest(http.MethodGet, "/greet?name=blocked", nil)
		recorder := httptest.NewRecorder()
		handler.NewHandler(service).ServeHTTP(recorder, request)

		if recorder.Code != http.StatusInternalServerError {
			t.Fatalf("status = %d, want %d", recorder.Code, http.StatusInternalServerError)
		}
	})
}
