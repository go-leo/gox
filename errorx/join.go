package errorx

import "errors"

func Join(err error, errs ...error) error {
	if err == nil {
		return errors.Join(errs...)
	}
	switch multiErr := err.(type) {
	case interface{ Unwrap() []error }:
		// standard join error
		return errors.Join(append(multiErr.Unwrap(), errs...)...)
	case interface{ WrappedErrors() []error }:
		// https://github.com/hashicorp/go-multierror/blob/main/multierror.go
		return errors.Join(append(multiErr.WrappedErrors(), errs...)...)
	case interface{ Errors() []error }:
		// https://github.com/uber-go/multierr/blob/master/error.go
		return errors.Join(append(multiErr.Errors(), errs...)...)
	default:
		return errors.Join(append([]error{err}, errs...)...)
	}
}

func Unwrap(err error) []error {
	switch multiErr := err.(type) {
	case interface{ Unwrap() []error }:
		// standard join error
		return multiErr.Unwrap()
	case interface{ WrappedErrors() []error }:
		// https://github.com/hashicorp/go-multierror/blob/main/multierror.go
		return multiErr.WrappedErrors()
	case interface{ Errors() []error }:
		// https://github.com/uber-go/multierr/blob/master/error.go
		return multiErr.Errors()
	default:
		return []error{err}
	}
}
