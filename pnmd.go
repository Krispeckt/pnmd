package pnmd

import (
	"io"
	"log/slog"

	h "github.com/TallSmaN/pnmd/internal/handler"
)

// Options defines configuration parameters for the pnmd log handler.
type Options struct {
	// Level specifies the minimum log level for output.
	Level slog.Level
	// CallerEnabled controls whether caller info is shown per log level.
	CallerEnabled map[slog.Level]bool
	// TimeFormat defines the timestamp format in log messages.
	TimeFormat string
	// Padding sets the left padding width for attributes and caller lines.
	Padding int
}

// GetLevel returns the configured minimum log level.
func (o *Options) GetLevel() slog.Level { return o.Level }

// IsCallerEnabled reports whether caller information is enabled for a given log level.
func (o *Options) IsCallerEnabled(l slog.Level) bool { return o.CallerEnabled[l] }

// GetPadding returns the configured padding width for log output.
func (o *Options) GetPadding() int { return o.Padding }

// GetTimeFormat returns the time format string used for log timestamps.
func (o *Options) GetTimeFormat() string { return o.TimeFormat }

// DefaultOptions returns a new Options instance initialized with default values.
func DefaultOptions() *Options {
	return &Options{
		Level: slog.LevelInfo,
		CallerEnabled: map[slog.Level]bool{
			slog.LevelDebug: true,
			slog.LevelInfo:  true,
			slog.LevelWarn:  true,
			slog.LevelError: true,
		},
		TimeFormat: "2006/01/02 15:04:05",
		Padding:    3,
	}
}

// NewHandler creates a new log handler using the provided options or defaults if nil.
func NewHandler(w io.Writer, opts *Options) *h.Handler {
	if opts == nil {
		opts = DefaultOptions()
	}

	return &h.Handler{
		Opts: opts,
		W:    w,
	}
}
