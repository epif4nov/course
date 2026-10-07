package handler

import (
	"context"
	"fmt"
	"net/http"
)

//go:generate go run go.uber.org/mock/mockgen@v0.6.0 -source=handler.go -destination=mocks/greeter_mock.go -package=mocks

type Greeter interface {
	Greet(context.Context, string) (string, error)
}

type Handler struct {
	service Greeter
}

func NewHandler(service Greeter) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		http.Error(writer, "name is required", http.StatusBadRequest)
		return
	}

	greeting, err := h.service.Greet(request.Context(), name)
	if err != nil {
		http.Error(writer, "service unavailable", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(writer, greeting)
}
