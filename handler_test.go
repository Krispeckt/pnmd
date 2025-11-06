package pnmd_test

import (
	"log/slog"
	"reflect"
	"testing"

	"github.com/TallSmaN/pnmd"
	h "github.com/TallSmaN/pnmd/internal/handler"
)

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
			},
		},
		{
			Name: "custom options",
			Opts: &pnmd.Options{
				Level: 1,
				CallerEnabled: map[slog.Level]bool{
					slog.LevelWarn: false,
				},
				Padding: 5,
			},
			Want: &h.Handler{
				Opts: &pnmd.Options{
					Level: 1,
					CallerEnabled: map[slog.Level]bool{
						slog.LevelWarn: false,
					},
					Padding: 5,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := pnmd.NewHandler(tt.Opts)

			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("NewHandler() = %+v, want %+v", got, tt.Want)
			}
		})
	}
}
