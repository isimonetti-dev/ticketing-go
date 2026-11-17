package main

import (
	"io"
	"log/slog"

	"github.com/isimonetti-dev/ticketing-go/internal/platform/config"
)

func newLogger(w io.Writer, format string) *slog.Logger {
	if format == config.LogFormatJSON {
		return slog.New(slog.NewJSONHandler(w, nil))
	}

	return slog.New(slog.NewTextHandler(w, nil))
}
