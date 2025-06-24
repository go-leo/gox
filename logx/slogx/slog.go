package slogx

import (
	"context"
	"log/slog"
)

// attrKey is a private type used as a key for storing attributes in the context.
type attrKey struct{}

// AppendContext appends attributes to the context.
func AppendContext(ctx context.Context, attrs ...slog.Attr) context.Context {
	src, _ := FromContext(ctx)
	dst := make([]slog.Attr, len(src), len(src)+len(attrs))
	copy(dst, src)
	dst = append(dst, attrs...)
	return context.WithValue(ctx, attrKey{}, dst)
}

// NewContext creates a new context with attributes.
func NewContext(ctx context.Context, attrs ...slog.Attr) context.Context {
	return context.WithValue(ctx, attrKey{}, attrs)
}

// FromContext returns the attributes from the context.
func FromContext(ctx context.Context) ([]slog.Attr, bool) {
	attr, ok := ctx.Value(attrKey{}).([]slog.Attr)
	return attr, ok
}

// WithContextHandler wraps a slog.Handler with context support.
func WithContextHandler(handler slog.Handler) slog.Handler {
	return &contextHandler{Handler: handler}
}

// contextHandler is a private type that implements slog.Handler and adds context support.
type contextHandler struct {
	slog.Handler
}

// Handle processes a log record and adds context attributes if present.
func (c contextHandler) Handle(ctx context.Context, record slog.Record) error {
	attrs, ok := FromContext(ctx)
	if ok {
		record.AddAttrs(attrs...)
	}
	return c.Handler.Handle(ctx, record)
}
