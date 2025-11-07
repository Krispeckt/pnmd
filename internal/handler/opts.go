package handler

import (
	"log/slog"
)

// Opts defines the interface for retrieving handler configuration settings.
type Opts interface {
	// GetLevel returns the minimum log level for output.
	GetLevel() slog.Level

	// IsCallerEnabled checks if caller information should be included for the given level.
	IsCallerEnabled(level slog.Level) bool

	// GetPadding returns the amount of padding to apply in log messages.
	GetPadding() int

	// GetTimeFormat returns the time format string used for log timestamps.
	GetTimeFormat() string
}
