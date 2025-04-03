package slogx_test

import (
	"context"
	"github.com/go-leo/gox/logx/slogx"
	"log/slog"
	"os"
	"testing"
	"time"
)

func TestAppendContext(t *testing.T) {
	ctx := context.Background()
	attr1 := slog.String("key1", "value1")
	attr2 := slog.Int("key2", 123)

	ctx = slogx.AppendContext(ctx, attr1)
	ctx = slogx.AppendContext(ctx, attr2)

	attrs, ok := slogx.FromContext(ctx)
	if !ok {
		t.Fatal("Failed to retrieve attributes from context")
	}

	if len(attrs) != 2 {
		t.Fatalf("Expected 2 attributes, got %d", len(attrs))
	}

	if attrs[0].Key != "key1" || attrs[0].Value.String() != "value1" {
		t.Errorf("Expected key1=value1, got %s=%v", attrs[0].Key, attrs[0].Value)
	}

	if attrs[1].Key != "key2" || attrs[1].Value.Int64() != 123 {
		t.Errorf("Expected key2=123, got %s=%v", attrs[1].Key, attrs[1].Value)
	}
}

func TestNewContext(t *testing.T) {
	ctx := context.Background()
	attr1 := slog.String("key1", "value1")
	attr2 := slog.Int("key2", 123)

	ctx = slogx.NewContext(ctx, attr1, attr2)

	attrs, ok := slogx.FromContext(ctx)
	if !ok {
		t.Fatal("Failed to retrieve attributes from context")
	}

	if len(attrs) != 2 {
		t.Fatalf("Expected 2 attributes, got %d", len(attrs))
	}

	if attrs[0].Key != "key1" || attrs[0].Value.String() != "value1" {
		t.Errorf("Expected key1=value1, got %s=%v", attrs[0].Key, attrs[0].Value)
	}

	if attrs[1].Key != "key2" || attrs[1].Value.Int64() != 123 {
		t.Errorf("Expected key2=123, got %s=%v", attrs[1].Key, attrs[1].Value)
	}
}

func TestFromContext(t *testing.T) {
	ctx := context.Background()
	attr1 := slog.String("key1", "value1")
	attr2 := slog.Int("key2", 123)

	ctx = slogx.NewContext(ctx, attr1, attr2)

	attrs, ok := slogx.FromContext(ctx)
	if !ok {
		t.Fatal("Failed to retrieve attributes from context")
	}

	if len(attrs) != 2 {
		t.Fatalf("Expected 2 attributes, got %d", len(attrs))
	}

	if attrs[0].Key != "key1" || attrs[0].Value.String() != "value1" {
		t.Errorf("Expected key1=value1, got %s=%v", attrs[0].Key, attrs[0].Value)
	}

	if attrs[1].Key != "key2" || attrs[1].Value.Int64() != 123 {
		t.Errorf("Expected key2=123, got %s=%v", attrs[1].Key, attrs[1].Value)
	}

	// Test with empty context
	ctx = context.Background()
	attrs, ok = slogx.FromContext(ctx)
	if ok {
		t.Fatal("Expected no attributes in empty context")
	}
}

func TestContextAttrHandler_Enabled(t *testing.T) {
	h := &slogx.ContextAttrHandler{
		Handler: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}),
	}

	ctx := context.Background()
	if !h.Enabled(ctx, slog.LevelDebug) {
		t.Error("Expected Enabled to return true for LevelDebug")
	}

	if h.Enabled(ctx, slog.LevelInfo) {
		t.Error("Expected Enabled to return false for LevelInfo")
	}
}

func TestContextAttrHandler_Handle(t *testing.T) {
	h := &slogx.ContextAttrHandler{
		Handler: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}),
	}

	ctx := context.Background()
	attr1 := slog.String("key1", "value1")
	ctx = slogx.AppendContext(ctx, attr1)

	record := slog.Record{
		Time:    time.Now(),
		Message: "this is message",
		Level:   slog.LevelDebug,
	}

	err := h.Handle(ctx, record)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestContextAttrHandler_WithAttrs(t *testing.T) {
	h := &slogx.ContextAttrHandler{
		Handler: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}),
	}

	attr1 := slog.String("key1", "value1")
	newHandler := h.WithAttrs([]slog.Attr{attr1})

	if _, ok := newHandler.(*slogx.ContextAttrHandler); !ok {
		t.Error("Expected newHandler to be of type *slogx.ContextAttrHandler")
	}
}

func TestContextAttrHandler_WithGroup(t *testing.T) {
	h := &slogx.ContextAttrHandler{
		Handler: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}),
	}

	newHandler := h.WithGroup("group")

	if _, ok := newHandler.(*slogx.ContextAttrHandler); !ok {
		t.Error("Expected newHandler to be of type *slogx.ContextAttrHandler")
	}
}
