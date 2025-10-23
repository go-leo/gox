package slogx

import (
	"context"
	"log/slog"
	"runtime"
)

// WithCallerSkipHandle wraps the given handler with a callerSkipHandler that skips the specified number of frames.
// This allows adjusting the reported caller location in log records.
func WithCallerSkipHandle(handler slog.Handler, skip int) slog.Handler {
	return &callerSkipHandler{Handler: handler, skip: skip}
}

// callerSkipHandler is a decorator for slog.Handler that modifies the program counter (PC)
// used for determining the caller location in log records.
type callerSkipHandler struct {
	slog.Handler     // Embedded Handler to delegate all other methods
	skip         int // Number of stack frames to skip when determining caller
}

// Handle processes a log record by setting its PC (program counter) field based on the current call stack,
// skipping the configured number of frames, then delegates to the wrapped handler.
func (h *callerSkipHandler) Handle(ctx context.Context, record slog.Record) error {
	var pcs [1]uintptr
	// Capture the program counter at the specified skip depth
	runtime.Callers(h.skip, pcs[:])
	// Set the record's PC to our captured value for accurate caller reporting
	record.PC = pcs[0]
	// Delegate handling to the wrapped handler
	return h.Handler.Handle(ctx, record)
}
