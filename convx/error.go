package convx

import "fmt"

var (
	failedCast    = "convx: failed to cast %#v of type %T to %T"
	failedCastErr = failedCast + ", %w"
)

func failedCastValue[E any](o any) (E, error) {
	var zero E
	return zero, fmt.Errorf(failedCast, o, o, zero)
}

func failedCastErrValue[E any](o any, err error) (E, error) {
	var zero E
	return zero, fmt.Errorf(failedCastErr, o, o, zero, err)
}
