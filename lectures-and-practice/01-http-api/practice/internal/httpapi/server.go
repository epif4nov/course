package httpapi

import (
	"encoding/json"
	"net/http"

	"http-api-practice/internal/api"
)

// Server implements the generated OpenAPI handlers with static in-memory data.
type Server struct{}

var items = []map[string]string{
	{"id": "1", "title": "The Go Programming Language"},
	{"id": "2", "title": "Learning HTTP"},
}

func (Server) Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (Server) ListItems(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, items)
}

func (Server) GetItem(w http.ResponseWriter, _ *http.Request, id string) {
	for _, item := range items {
		if item["id"] == id {
			writeJSON(w, http.StatusOK, item)
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"message": "item not found"})
}

var _ api.ServerInterface = Server{}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
