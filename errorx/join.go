package errorx

import "errors"

func Join(err error, errs ...error) error {
	// standard join error
	if joinError, ok := err.(interface{ Unwrap() []error }); ok {
		return errors.Join(append(joinError.Unwrap(), errs...)...)
	}
	// https://github.com/hashicorp/go-multierror/blob/main/multierror.go
	if multiError, ok := err.(interface{ WrappedErrors() []error }); ok {
		return errors.Join(append(multiError.WrappedErrors(), errs...)...)
	}
	// https://github.com/uber-go/multierr/blob/master/error.go
	if multiErr, ok := err.(interface{ Errors() []error }); ok {
		return errors.Join(append(multiErr.Errors(), errs...)...)
	}
	return errors.Join(append([]error{err}, errs...)...)
}
