// Package slogx 提供了对 slog 日志库的扩展功能，主要增加了格式化日志记录的能力
package slogx

import (
	"context"     // 用于传递上下文信息
	"fmt"         // 用于字符串格式化
	"log/slog"    // Go 标准库的日志包
	"sync/atomic" // 用于原子操作，保证并发安全
)

// formatLogger 是一个全局的原子指针，指向 slog.Logger 实例
// 使用 atomic.Pointer 保证在并发环境下的安全访问
var formatLogger atomic.Pointer[slog.Logger]

// SetFormatLogger 设置全局格式化日志记录器
// 该函数会包装传入的 logger，添加调用者信息跳过和上下文处理功能
// 参数 l: 要设置的 slog.Logger 实例
func SetFormatLogger(l *slog.Logger) {
	// 使用 WithCallerSkipHandle 跳过 5 层调用栈，确保显示正确的调用者位置
	// 使用 WithContextHandler 处理上下文信息
	formatLogger.Store(slog.New(WithCallerSkip(WithContext(l.Handler()), 5)))
}

// Debugf 记录 DEBUG 级别的格式化日志
// 参数 ctx: 上下文信息
// 参数 format: 格式化字符串
// 参数 a: 要插入到格式化字符串中的参数
func Debugf(ctx context.Context, format string, a ...any) {
	// 加载当前的日志记录器并记录 DEBUG 级别日志
	// 使用 fmt.Sprintf 将格式化字符串和参数组合成最终消息
	formatLogger.Load().DebugContext(ctx, fmt.Sprintf(format, a...))
}

// Infof 记录 INFO 级别的格式化日志
// 参数 ctx: 上下文信息
// 参数 format: 格式化字符串
// 参数 a: 要插入到格式化字符串中的参数
func Infof(ctx context.Context, format string, a ...any) {
	// 加载当前的日志记录器并记录 INFO 级别日志
	formatLogger.Load().InfoContext(ctx, fmt.Sprintf(format, a...))
}

// Warnf 记录 WARN 级别的格式化日志
// 参数 ctx: 上下文信息
// 参数 format: 格式化字符串
// 参数 a: 要插入到格式化字符串中的参数
func Warnf(ctx context.Context, format string, a ...any) {
	// 加载当前的日志记录器并记录 WARN 级别日志
	formatLogger.Load().WarnContext(ctx, fmt.Sprintf(format, a...))
}

// Errorf 记录 ERROR 级别的格式化日志
// 参数 ctx: 上下文信息
// 参数 format: 格式化字符串
// 参数 a: 要插入到格式化字符串中的参数
func Errorf(ctx context.Context, format string, a ...any) {
	// 加载当前的日志记录器并记录 ERROR 级别日志
	formatLogger.Load().ErrorContext(ctx, fmt.Sprintf(format, a...))
}

// Logf 记录指定级别的格式化日志
// 参数 ctx: 上下文信息
// 参数 level: 日志级别 (slog.Level)
// 参数 format: 格式化字符串
// 参数 a: 要插入到格式化字符串中的参数
func Logf(ctx context.Context, level slog.Level, format string, a ...any) {
	// 加载当前的日志记录器并记录指定级别日志
	formatLogger.Load().Log(ctx, level, fmt.Sprintf(format, a...))
}
