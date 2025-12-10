// Package slogx 提供了对 slog 日志库的扩展功能
package slogx

import (
	"context"  // 用于处理日志上下文
	"log/slog" // Go标准库的日志包
	"runtime"  // 用于获取程序计数器(PC)和调用栈信息
)

// WithCallerSkip 包装给定的处理器，创建一个 callerSkipHandler，
// 该处理器会在确定调用者位置时跳过指定数量的调用栈帧数。
// 这允许调整日志记录中报告的调用者位置，确保显示正确的源代码位置。
// 参数:
//
//	handler: 要包装的基础 slog.Handler
//	skip: 要跳过的调用栈帧数
//
// 返回值:
//
//	返回包装后的 slog.Handler
func WithCallerSkip(handler slog.Handler, skip int) slog.Handler {
	return &callerSkipHandler{Handler: handler, skip: skip}
}

// callerSkipHandler 是 slog.Handler 的装饰器，它修改用于确定日志记录中调用者位置的程序计数器(PC)
// 通过跳过指定的调用栈帧数来实现更准确的调用者位置报告
type callerSkipHandler struct {
	slog.Handler     // 嵌入的基础 Handler，用于委托处理其他方法
	skip         int // 在确定调用者位置时要跳过的调用栈帧数
}

// Handle 处理日志记录，通过设置记录的PC(程序计数器)字段来基于当前调用栈确定调用者位置，
// 跳过配置的帧数后，再委托给包装的处理器进行处理。
// 参数:
//
//	ctx: 日志上下文
//	record: 要处理的日志记录
//
// 返回值:
//
//	处理过程中可能出现的错误
func (h *callerSkipHandler) Handle(ctx context.Context, record slog.Record) error {
	// 创建一个长度为1的uintptr数组来存储程序计数器
	var pcs [1]uintptr

	// 使用 runtime.Callers 捕获指定跳过深度的程序计数器
	// h.skip 参数指定了要跳过的调用栈帧数，这样可以获得更准确的调用者位置
	runtime.Callers(h.skip, pcs[:])

	// 将捕获到的程序计数器设置到日志记录的PC字段中
	// 这样可以确保日志输出中显示的调用者位置是准确的
	record.PC = pcs[0]

	// 将处理后的记录委托给包装的基础处理器进行实际处理
	return h.Handler.Handle(ctx, record)
}
