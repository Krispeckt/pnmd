package opts

import "log/slog"

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

func Default() *Options {
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
