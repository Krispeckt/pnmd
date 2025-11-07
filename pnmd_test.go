package pnmd_test

import (
	"log/slog"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/TallSmaN/pnmd"
	h "github.com/TallSmaN/pnmd/internal/handler"
)

func TestExample(t *testing.T) {
	logger := slog.New(
		pnmd.NewHandler(os.Stdout, &pnmd.Options{
			Level: slog.LevelDebug,
			CallerEnabled: map[slog.Level]bool{
				slog.LevelDebug: true,
				slog.LevelInfo:  true,
				slog.LevelWarn:  true,
				slog.LevelError: true,
			},
			TimeFormat: "2006/01/02 15:04:05",
			Padding:    3,
		}),
	)

	logger.Debug("initializing cache subsystem", "cache", "redis", "host", "localhost", "port", 6379) //nolint

	logger.Info("cache connected") //nolint

	logger.Warn("slow query detected", "duration_ms", 1823, "query", "SELECT * FROM users WHERE active=1", "user", "analytics-worker") //nolint

	logger.Error("failed to write audit event", "error", "disk full", "path", "/var/log/audit.json", "component", "audit", "retry_in_sec", 30) //nolint

	logger.Debug("config reloaded", "file", "/etc/app/config.yaml", "changes", 5) //nolint

	logger.Info("http server started", "addr", ":8080", "threads", 8) //nolint

	logger.Warn("deprecated API usage", "endpoint", "/v1/legacy", "client", "mobile-android", "version", "1.2.0") //nolint

	logger.Error("user authentication failed", "user", "john", "ip", "192.168.1.42", "reason", "invalid token") //nolint

	logger.Debug("background job finished", "job_id", "import-2025-11-05", "rows", 152_000, "duration_sec", 94.2) //nolint

	logger.Info("graceful shutdown complete", "uptime_min", 238) //nolint

	t.Skip()
}

func TestOptions(t *testing.T) {
	tests := []struct {
		Name string
		Opts *pnmd.Options
		Want *h.Handler
	}{
		{
			Name: "nil options -> default options",
			Opts: nil,
			Want: &h.Handler{
				Opts: pnmd.DefaultOptions(),
				W:    os.Stdout,
			},
		},
		{
			Name: "custom options",
			Opts: &pnmd.Options{
				Level: slog.LevelInfo,
				CallerEnabled: map[slog.Level]bool{
					slog.LevelWarn: false,
				},
				TimeFormat: time.TimeOnly,
				Padding:    5,
			},
			Want: &h.Handler{
				Opts: &pnmd.Options{
					Level: slog.LevelInfo,
					CallerEnabled: map[slog.Level]bool{
						slog.LevelWarn: false,
					},
					TimeFormat: time.TimeOnly,
					Padding:    5,
				},
				W: os.Stdout,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := pnmd.NewHandler(os.Stdout, tt.Opts)

			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("NewHandler() = %+v, want %+v", got, tt.Want)
			}
		})
	}
}
