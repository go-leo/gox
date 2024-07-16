// Package goworkers
// dpaks/goworkers提供了更便利的Submi方法提交任务以及Worker数、任务数等查
// 询方法、关闭Pool的方法。它的任务的执行结果需要在ResultChan和ErrChan中去获取，没有提供阻塞
// 的方法，但是它可以在初始化的时候设置Worker的数量和任务数。
// See: https://github.com/dpaks/goworkers
package goworkers
