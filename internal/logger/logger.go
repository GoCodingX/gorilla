package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func init() { //nolint:gochecknoinits
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}

func ErrorAndExit(message string, err error) {
	logger.Error(message, slog.String("error", err.Error()))
	os.Exit(1)
}
