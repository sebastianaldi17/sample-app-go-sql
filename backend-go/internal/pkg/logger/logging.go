package logger

import (
	"os"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrslog"
	"github.com/newrelic/go-agent/v3/newrelic"

	"log/slog"
)

var (
	logger *slog.Logger
)

func InitLogging(app *newrelic.Application) {
	if app != nil {
		instrumentedTextHandler := nrslog.TextHandler(app, os.Stdout, &slog.HandlerOptions{})
		logger = slog.New(instrumentedTextHandler)
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	}
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	logger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}
