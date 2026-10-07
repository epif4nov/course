package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	" http-api-practice/internal/api"
	" http-api-practice/internal/httpapi"
	" http-api-practice/internal/logging"
)

// New wires the generated routes, handlers, and request logging middleware.
func New(logger *log.Logger) http.Handler {
	router := chi.NewRouter()
	router.Use(logging.Middleware(logger))
	return api.HandlerFromMux(httpapi.Server{}, router)
}
