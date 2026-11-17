package httpserver_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/isimonetti-dev/ticketing-go/internal/platform/httpserver"
)

func TestHealth(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	httpserver.Health(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want %q", got, "application/json")
	}

	var body struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("decode body: %v", err)
	}

	if body.Status != "ok" {
		t.Errorf("status field = %q, want %q", body.Status, "ok")
	}
}
