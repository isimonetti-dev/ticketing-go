package config_test

import (
	"testing"
	"time"

	"github.com/isimonetti-dev/ticketing-go/internal/platform/config"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		want    config.Config
		wantErr bool
	}{
		{
			name: "defaults",
			want: config.Config{HTTPAddr: ":8080", ShutdownTimeout: 20 * time.Second, LogFormat: "text"},
		},
		{
			name: "custom http addr",
			env:  map[string]string{"HTTP_ADDR": "127.0.0.1:9000"},
			want: config.Config{HTTPAddr: "127.0.0.1:9000", ShutdownTimeout: 20 * time.Second, LogFormat: "text"},
		},
		{
			name: "custom shutdown timeout",
			env:  map[string]string{"SHUTDOWN_TIMEOUT": "1ns"},
			want: config.Config{HTTPAddr: ":8080", ShutdownTimeout: time.Nanosecond, LogFormat: "text"},
		},
		{
			name:    "invalid shutdown timeout",
			env:     map[string]string{"SHUTDOWN_TIMEOUT": "soon"},
			wantErr: true,
		},
		{
			name:    "zero shutdown timeout",
			env:     map[string]string{"SHUTDOWN_TIMEOUT": "0s"},
			wantErr: true,
		},
		{
			name:    "negative shutdown timeout",
			env:     map[string]string{"SHUTDOWN_TIMEOUT": "-5s"},
			wantErr: true,
		},
		{
			name: "text log format",
			env:  map[string]string{"LOG_FORMAT": "text"},
			want: config.Config{HTTPAddr: ":8080", ShutdownTimeout: 20 * time.Second, LogFormat: "text"},
		},
		{
			name: "json log format",
			env:  map[string]string{"LOG_FORMAT": "json"},
			want: config.Config{HTTPAddr: ":8080", ShutdownTimeout: 20 * time.Second, LogFormat: "json"},
		},
		{
			name:    "invalid log format",
			env:     map[string]string{"LOG_FORMAT": "xml"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range []string{"HTTP_ADDR", "SHUTDOWN_TIMEOUT", "LOG_FORMAT"} {
				t.Setenv(key, tt.env[key])
			}

			got, err := config.Load()

			if tt.wantErr {
				if err == nil {
					t.Fatalf("Load() error = nil, want error (config = %+v)", got)
				}
				return
			}

			if err != nil {
				t.Fatalf("Load() error = %v, want nil", err)
			}

			if got != tt.want {
				t.Errorf("Load() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
