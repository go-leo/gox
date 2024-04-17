package brave

import (
	"fmt"
	"log"
)

func Do(f func(), rs ...func(p any)) {
	defer func() {
		if len(rs) <= 0 {
			rs = append(rs, func(p any) {
				log.Printf("panic triggered: %v\n", p)
			})
		}
		r := rs[0]
		if r == nil {
			return
		}
		if p := recover(); p != nil {
			r(p)
		}
	}()
	f()
}

func Go(f func(), rs ...func(p any)) {
	go Do(f, rs...)
}

func DoE(f func() error, rs ...func(p any) error) (err error) {
	defer func() {
		if len(rs) <= 0 {
			rs = append(rs, func(p any) error {
				return fmt.Errorf("panic triggered: %v", p)
			})
		}
		r := rs[0]
		if r == nil {
			return
		}
		if p := recover(); p != nil {
			err = r(p)
		}
	}()
	return f()
}

// GoE If the execution fails, error is sent to errC, and errC is closed,
// if the execution succeeds, errC is closed directly.
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

func DoRE[R any](f func() (R, error), rs ...func(p any) error) (ret R, err error) {
	defer func() {
		if len(rs) <= 0 {
			rs = append(rs, func(p any) error {
				return fmt.Errorf("panic triggered: %v", p)
			})
		}
		r := rs[0]
		if r == nil {
			return
		}
		if p := recover(); p != nil {
			err = r(p)
		}
	}()
	return f()
}

// GoRE If execution fails, retC is closed directly, error is sent to errC and errC is closed,
// if execution succeeds, errC is closed and the result is sent to retC, and retC is closed.
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
