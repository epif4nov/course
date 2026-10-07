package app

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoutesAndRequestLogging(t *testing.T) {
	var logs bytes.Buffer
	handler := New(log.New(&logs, "", 0))

	tests := []struct {
		name       string
		path       string
		wantStatus int
		wantBody   string
	}{
		{name: "health", path: "/health", wantStatus: http.StatusOK, wantBody: `"status":"ok"`},
		{name: "list items", path: "/items", wantStatus: http.StatusOK, wantBody: `"title":"The Go Programming Language"`},
		{name: "get item", path: "/items/2", wantStatus: http.StatusOK, wantBody: `"title":"Learning HTTP"`},
		{name: "missing item", path: "/items/missing", wantStatus: http.StatusNotFound, wantBody: `"message":"item not found"`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, test.path, nil)
			response := httptest.NewRecorder()
			handler.ServeHTTP(response, request)

			if response.Code != test.wantStatus {
				t.Fatalf("status = %d, want %d", response.Code, test.wantStatus)
			}
			if !strings.Contains(response.Body.String(), test.wantBody) {
				t.Fatalf("body %q does not contain %q", response.Body.String(), test.wantBody)
			}
			if !strings.Contains(logs.String(), "method=GET uri="+test.path+" status=") {
				t.Fatalf("request log does not contain method, URI, and status: %q", logs.String())
			}
		})
	}
}
