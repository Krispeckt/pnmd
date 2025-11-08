package pnmd

import (
	"io"

	h "github.com/TallSmaN/pnmd/internal/handler"
	"github.com/TallSmaN/pnmd/internal/opts"
)

type Options = opts.Options

// DefaultOptions returns a new Options instance initialized with default values.
func DefaultOptions() *Options {
	return opts.Default()
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
