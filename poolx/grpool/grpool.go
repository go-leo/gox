// Package grpool
// grpool创建Pool的时候需要提供Worker的数量和等待执行的任务的最大数量，任务的提
// 交是直接往Channel放入任务。
// See: https://github.com/ivpusic/grpool
package grpool