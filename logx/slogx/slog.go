package slogx

import (
	"context"
	"log/slog"
)

type attrKey struct{}

func AppendContext(ctx context.Context, attrs ...slog.Attr) context.Context {
	src, _ := FromContext(ctx)
	dst := make([]slog.Attr, len(src), len(src)+len(attrs))
	copy(dst, src)
	dst = append(dst, attrs...)
	return context.WithValue(ctx, attrKey{}, dst)
}

func NewContext(ctx context.Context, attrs ...slog.Attr) context.Context {
	return context.WithValue(ctx, attrKey{}, attrs)
}

func FromContext(ctx context.Context) ([]slog.Attr, bool) {
	attr, ok := ctx.Value(attrKey{}).([]slog.Attr)
	return attr, ok
}

func WithContextHandler(handler slog.Handler) slog.Handler {
	return &contextHandler{Handler: handler}
}

type contextHandler struct {
	slog.Handler
}

func (c contextHandler) Handle(ctx context.Context, record slog.Record) error {
	attrs, ok := FromContext(ctx)
	if ok {
		record.AddAttrs(attrs...)
	}
	return c.Handler.Handle(ctx, record)
}
