package handler

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

type Handler struct {
	Opts Opts
}

// Enabled reports whether the given level is enabled by cfg.Level.
func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.Opts.GetLevel()
}

// Handle formats the record as a  and writes it to stdout.
// Output includes timestamp, level, message, attrs, and optional caller info.
// The context is ignored.
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	var builder strings.Builder

	builder.WriteString(pterm.Gray(time.Now().Format("2006-01-02 15:04:05")))
	builder.WriteString(" ")

	style := styleForLevel(r.Level)
	builder.WriteString(style.Sprintf("%s", r.Level.String()[:4]))
	builder.WriteString(" ")

	builder.WriteString(r.Message)

	var arguments []string
	r.Attrs(func(attr slog.Attr) bool {
		arguments = append(arguments, style.Sprintf("%s: ", attr.Key)+fmt.Sprint(attr.Value))
		return true
	})

	if enabled := h.Opts.IsCallerEnabled(r.Level); enabled {
		if fn := runtime.FuncForPC(r.PC); fn != nil {
			file, line := fn.FileLine(r.PC)
			arguments = append(arguments, pterm.NewStyle(pterm.Bold, pterm.FgGray).Sprintf("caller: %s:%d", file, line))
		}
	}

	padding := len(time.Time{}.Format("2006/01/02 15:04:05")) + h.Opts.GetPadding()
	for i, arg := range arguments {
		pipe := "└"
		if i < len(arguments)-1 {
			pipe = "├"
		}
		builder.WriteString("\n")
		builder.WriteString(strings.Repeat(" ", padding))
		builder.WriteString(pipe)
		builder.WriteString(" ")
		builder.WriteString(arg)
	}

	fmt.Println(builder.String())
	return nil
}

// WithAttrs returns h unchanged
func (h *Handler) WithAttrs(_ []slog.Attr) slog.Handler { return h }

// WithGroup returns h unchanged
func (h *Handler) WithGroup(_ string) slog.Handler { return h }
