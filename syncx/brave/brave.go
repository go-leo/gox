package brave

import (
	"fmt"
)

// Do 用于执行传入的函数 f，并在 f 执行过程中捕获任何可能发生的 panic。
// 如果发生 panic，可以选择性地调用传入的恢复函数 rs 中的第一个函数来处理 panic，
// 如果没有提供恢复函数，则默认记录 panic 信息。
func Do(f func(), rs ...func(p any)) {
	defer func() {
		if p := recover(); p != nil {
			var r func(p any)
			if len(rs) > 0 {
				r = rs[0]
			}
			if r == nil {
				r = func(p any) {
					fmt.Printf("brave: panic triggered: %v\n", p)
				}
			}
			r(p)
		}
	}()
	f()
}

// DoE 和 Do 功能上一致，但返回一个错误信息。
func DoE(f func() error, rs ...func(p any) error) (err error) {
	defer func() {
		if p := recover(); p != nil {
			var r func(p any) error
			if len(rs) > 0 {
				r = rs[0]
			}
			if r == nil {
				r = func(p any) error {
					return fmt.Errorf("brave: panic triggered: %v", p)
				}
			}
			err = r(p)
		}
	}()
	return f()
}

// DoRE 和 DoE 功能一致，但返回一个结果和错误信息。
func DoRE[R any](f func() (R, error), rs ...func(p any) error) (_ R, err error) {
	defer func() {
		if p := recover(); p != nil {
			var r func(p any) error
			if len(rs) > 0 {
				r = rs[0]
			}
			if r == nil {
				r = func(p any) error {
					return fmt.Errorf("brave: panic triggered: %v", p)
				}
			}
			err = r(p)
		}
	}()
	return f()
}

// Go 异步执行 Do 函数。
func Go(f func(), rs ...func(p any)) {
	go Do(f, rs...)
}

// GoE 异步执行 DoE 函数。返回一个错误通道。
func GoE(f func() error, rs ...func(p any) error) <-chan error {
	errC := make(chan error, 1)
	go func() {
		defer close(errC)
		err := DoE(f, rs...)
		if err != nil {
			errC <- err
		}
	}()
	return errC
}

// GoRE 异步执行 DoRE 函数。返回一个结果通道和一个错误通道。
func GoRE[R any](f func() (R, error), rs ...func(p any) error) (<-chan R, <-chan error) {
	retC := make(chan R, 1)
	errC := make(chan error, 1)
	go func() {
		defer close(errC)
		defer close(retC)
		ret, err := DoRE[R](f, rs...)
		if err != nil {
			errC <- err
			return
		}
		retC <- ret
	}()
	return retC, errC
}
