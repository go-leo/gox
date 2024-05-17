package strconvx

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

func FormatBool(b bool) string {
	return strconv.FormatBool(b)
}

func FormatUint[Unsigned constraints.Unsigned](i Unsigned, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

func FormatInt[Signed constraints.Signed](i Signed, base int) string {
	return strconv.FormatInt(int64(i), base)
}

func FormatFloat[Float constraints.Float](f Float, fmt byte, prec, bitSize int) string {
	return strconv.FormatFloat(float64(f), fmt, prec, bitSize)
}

func FormatBoolSlice(s []bool) []string {
	if s == nil {
		return nil
	}
	r := make([]string, 0, len(s))
	for _, b := range s {
		r = append(r, FormatBool(b))
	}
	return r
}

func FormatUintSlice[Unsigned constraints.Unsigned](s []Unsigned, base int) []string {
	if s == nil {
		return nil
	}
	r := make([]string, 0, len(s))
	for _, i := range s {
		r = append(r, FormatUint(i, base))
	}
	return r
}

func FormatIntSlice[Signed constraints.Signed](s []Signed, base int) []string {
	if s == nil {
		return nil
	}
	r := make([]string, 0, len(s))
	for _, i := range s {
		r = append(r, FormatInt(i, base))
	}
	return r
}

func FormatFloatSlice[Float constraints.Float](s []Float, fmt byte, prec, bitSize int) []string {
	if s == nil {
		return nil
	}
	r := make([]string, 0, len(s))
	for _, f := range s {
		r = append(r, FormatFloat(float64(f), fmt, prec, bitSize))
	}
	return r
}
