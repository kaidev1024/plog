package plog

import (
	"log/slog"
	"os"
	"runtime"
)

// Logger instance (global)
var logger *slog.Logger

// InitLogger initializes the structured logger
func InitLogger() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger.Info("plog initialized")
}

// LogError logs an error with structured logging (package, function, line number, and message)
func LogError(err error) {
	if err == nil {
		return
	}

	// Get caller information
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		logger.Error("Failed to retrieve caller info", "error", err)
		return
	}

	// Get function name
	funcName := runtime.FuncForPC(pc).Name()

	// Log the error with structured fields
	logger.Error("An error occurred--------------------",
		slog.String("error", err.Error()),
		slog.String("file", file),
		slog.String("function", funcName),
		slog.Int("line", line),
	)
}

func PrintSlice(arr []string) {
	for i, id := range arr {
		println("++++++-----", i, id)
	}
}
