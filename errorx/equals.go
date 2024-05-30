package errorx

import (
	"errors"
	"reflect"
)

// Equals 判断两个错误是否相等
func Equals(err, target error) bool {
	if err == nil && target == nil {
		return true
	}
	if err != nil && target != nil {
		return equals(err, target)
	}
	return false
}

func equals(err error, target error) bool {
	if errors.Is(err, target) {
		return true
	}
	if reflect.TypeOf(err) != reflect.TypeOf(target) {
		return false
	}
	return err.Error() == err.Error()
}
