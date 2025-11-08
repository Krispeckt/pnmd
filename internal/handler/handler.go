package handler

import (
	"context"
	"io"
	"log/slog"

	"github.com/TallSmaN/pnmd/internal/builder"
	"github.com/TallSmaN/pnmd/internal/opts"
)

// Handler implements slog.Handler and writes formatted log records.
type Handler struct {
	// Opts defines configuration options for log formatting and behavior.
	Opts *opts.Options
	// W is the output destination for log entries.
	W io.Writer
}

// Enabled reports whether the given level is enabled by cfg.Level.
func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.Opts.Level
}

// Handle formats the record as a  and writes it to stdout.
// Output includes timestamp, level, message, attrs, and optional caller info.
// The context is ignored.
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	b := builder.NewBuilder(h.Opts, &r)

	b.WriteTime()
	b.WriteLevel()
	b.WriteMessage()
	b.WriteAttrs()

	_, _ = h.W.Write([]byte(b.Build()))
	return nil
}

// WithAttrs returns h unchanged
func (h *Handler) WithAttrs(_ []slog.Attr) slog.Handler { return h }

// WithGroup returns h unchanged
func (h *Handler) WithGroup(_ string) slog.Handler { return h }
