package convx

import "errors"

var (
	ErrValueIsNULL        = errors.New("convx: unable to convert NULL value")
	ErrNegativeNotAllowed = errors.New("convx: unable to convert negative value")
)
