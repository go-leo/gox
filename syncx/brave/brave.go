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

func GoE(f func() error, rs ...func(p any) error) <-chan error {
	errC := make(chan error)
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

func GoRE[R any](f func() (R, error), rs ...func(p any) error) (<-chan R, <-chan error) {
	retC := make(chan R)
	errC := make(chan error)
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
