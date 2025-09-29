package logging

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
)

func NewLogger(w *os.File) *slog.Logger {
	logger := slog.New(
		tint.NewHandler(w, &tint.Options{
			NoColor:    !isatty.IsTerminal(w.Fd()),
			TimeFormat: time.RFC3339Nano,
		}),
	)
	return logger
}
