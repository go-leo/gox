package brave

func Do(f func(), r func(p any)) {
	defer func() {
		// if r is nil, which means panics are not recovered.
		if r == nil {
			return
		}
		if p := recover(); p != nil {
			r(p)
		}
	}()
	f()
}

func DoE(f func() error, r func(p any) error) (err error) {
	defer func() {
		// if r is nil, which means panics are not recovered.
		if r == nil {
			return
		}
		if p := recover(); p != nil {
			err = r(p)
		}
	}()
	return f()
}

func DoRE(f func() (any, error), r func(p any) error) (ret any, err error) {
	defer func() {
		// if r is nil, which means panics are not recovered.
		if r == nil {
			return
		}
		if p := recover(); p != nil {
			err = r(p)
		}
	}()
	return f()
}

func Go(f func(), r func(p any)) {
	go Do(f, r)
}

func GoE(f func() error, r func(p any) error) <-chan error {
	errC := make(chan error)
	go func() {
		defer close(errC)
		err := DoE(f, r)
		if err != nil {
			errC <- err
		}
	}()
	return errC
}

func GoRE(f func() (any, error), r func(p any) error) (<-chan any, <-chan error) {
	retC := make(chan any)
	errC := make(chan error)
	go func() {
		defer close(errC)
		defer close(retC)
		ret, err := DoRE(f, r)
		if err != nil {
			errC <- err
			return
		}
		retC <- ret
	}()
	return retC, errC
}
