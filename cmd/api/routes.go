package main

import (
	"net/http"

	"github.com/isimonetti-dev/ticketing-go/internal/platform/httpserver"
)

func newRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", httpserver.Health)

	return mux
}
