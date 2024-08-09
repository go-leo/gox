// Package workerpool
// gammazero/workerpool可以无限制地提交任务，提供了更便利的Submit和SubmitWait方法提交任务，
// 还可以提供当前的worker数和任务数以及关闭Pool的功能。
// See: https://github.com/gammazero/workerpool
package workerpool

import "github.com/gammazero/workerpool"

type Gopher struct {
	Pool *workerpool.WorkerPool
}

func (g Gopher) Go(f func()) error {
	g.Pool.Submit(f)
	return nil
}
