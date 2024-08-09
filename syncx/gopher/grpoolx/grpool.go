// Package grpoolx
// grpool创建Pool的时候需要提供Worker的数量和等待执行的任务的最大数量，任务的提
// 交是直接往Channel放入任务。
package grpoolx

import "github.com/ivpusic/grpool"

type Gopher struct {
	Pool *grpool.Pool
}

func (g Gopher) Go(f func()) error {
	g.Pool.JobQueue <- f
	return nil
}
