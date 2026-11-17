package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		method     string
		path       string
		wantStatus int
	}{
		{name: "get healthz", method: http.MethodGet, path: "/healthz", wantStatus: http.StatusOK},
		{name: "post healthz", method: http.MethodPost, path: "/healthz", wantStatus: http.StatusMethodNotAllowed},
		{name: "unknown route", method: http.MethodGet, path: "/unknown", wantStatus: http.StatusNotFound},
	}

	router := newRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequestWithContext(t.Context(), tt.method, tt.path, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("%s %s status = %d, want %d", tt.method, tt.path, rec.Code, tt.wantStatus)
			}
		})
	}
}
