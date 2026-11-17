package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/isimonetti-dev/ticketing-go/internal/platform/config"
	"github.com/isimonetti-dev/ticketing-go/internal/platform/httpserver"
)

func main() {
	if err := run(context.Background()); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	logger := newLogger(os.Stdout, cfg.LogFormat)
	srv := httpserver.New(newRouter())

	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", cfg.HTTPAddr)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	serveErrc := make(chan error, 1)
	go func() {
		serveErrc <- srv.Serve(listener)
	}()

	logger.Info("server started", "addr", listener.Addr().String())

	select {
	case <-ctx.Done():
		stop()
		logger.Info("shutdown requested")
	case err := <-serveErrc:
		return fmt.Errorf("serve: %w", err)
	}

	shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), cfg.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		_ = srv.Close()
		return fmt.Errorf("shutdown: %w", err)
	}

	if err := <-serveErrc; !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("serve: %w", err)
	}

	logger.Info("server stopped")

	return nil
}
