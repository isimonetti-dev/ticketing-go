package config

import (
	"fmt"
	"os"
	"time"
)

const (
	envHTTPAddr        = "HTTP_ADDR"
	envShutdownTimeout = "SHUTDOWN_TIMEOUT"
	envLogFormat       = "LOG_FORMAT"

	defaultHTTPAddr        = ":8080"
	defaultShutdownTimeout = 20 * time.Second
	defaultLogFormat       = LogFormatText
)

const (
	LogFormatText = "text"
	LogFormatJSON = "json"
)

type Config struct {
	HTTPAddr        string
	ShutdownTimeout time.Duration
	LogFormat       string
}

func Load() (Config, error) {
	cfg := Config{
		HTTPAddr:        defaultHTTPAddr,
		ShutdownTimeout: defaultShutdownTimeout,
		LogFormat:       defaultLogFormat,
	}

	if addr := os.Getenv(envHTTPAddr); addr != "" {
		cfg.HTTPAddr = addr
	}

	if raw := os.Getenv(envShutdownTimeout); raw != "" {
		timeout, err := time.ParseDuration(raw)
		if err != nil {
			return Config{}, fmt.Errorf("invalid %s: %w", envShutdownTimeout, err)
		}

		if timeout <= 0 {
			return Config{}, fmt.Errorf("%s must be positive, got %s", envShutdownTimeout, raw)
		}

		cfg.ShutdownTimeout = timeout
	}

	if format := os.Getenv(envLogFormat); format != "" {
		if format != LogFormatText && format != LogFormatJSON {
			return Config{}, fmt.Errorf(
				"%s must be %q or %q, got %q",
				envLogFormat,
				LogFormatText,
				LogFormatJSON,
				format,
			)
		}

		cfg.LogFormat = format
	}

	return cfg, nil
}
