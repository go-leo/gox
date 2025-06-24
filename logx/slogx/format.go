package slogx

import (
	"context"
	"fmt"
	"log/slog"
)

// Debugf formats and logs a debug message using slog.
func Debugf(format string, a ...any) {
	slog.Debug(fmt.Sprintf(format, a...))
}

// DebugContextf formats and logs a debug message with context using slog.
func DebugContextf(ctx context.Context, format string, a ...any) {
	slog.DebugContext(ctx, fmt.Sprintf(format, a...))
}

// Infof formats and logs an info message using slog.
func Infof(format string, a ...any) {
	slog.Info(fmt.Sprintf(format, a...))
}

// InfoContextf formats and logs an info message with context using slog.
func InfoContextf(ctx context.Context, format string, a ...any) {
	slog.InfoContext(ctx, fmt.Sprintf(format, a...))
}

// Warnf formats and logs a warning message using slog.
func Warnf(format string, a ...any) {
	slog.Warn(fmt.Sprintf(format, a...))
}

// WarnContextf formats and logs a warning message with context using slog.
func WarnContextf(ctx context.Context, format string, a ...any) {
	slog.WarnContext(ctx, fmt.Sprintf(format, a...))
}

// Errorf formats and logs an error message using slog.
func Errorf(format string, a ...any) {
	slog.Error(fmt.Sprintf(format, a...))
}

// ErrorContextf formats and logs an error message with context using slog.
func ErrorContextf(ctx context.Context, format string, a ...any) {
	slog.ErrorContext(ctx, fmt.Sprintf(format, a...))
}

// Logf formats and logs a message at a specific level using slog.
func Logf(ctx context.Context, level slog.Level, format string, a ...any) {
	slog.Log(ctx, level, fmt.Sprintf(format, a...))
}
