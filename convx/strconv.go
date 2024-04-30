package convx

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

func ParseBoolSlice(s []string) ([]bool, error) {
	r := make([]bool, 0, len(s))
	for _, str := range s {
		b, err := strconv.ParseBool(str)
		if err != nil {
			return nil, err
		}
		r = append(r, b)
	}
	return r, nil
}

func ParseIntSlice[Integer constraints.Signed](s []string, base int, bitSize int) ([]Integer, error) {
	r := make([]Integer, 0, len(s))
	for _, str := range s {
		i, err := strconv.ParseInt(str, base, bitSize)
		if err != nil {
			return nil, err
		}
		r = append(r, Integer(i))
	}
	return r, nil
}

func ParseUintSlice[Integer constraints.Unsigned](s []string, base int, bitSize int) ([]Integer, error) {
	r := make([]Integer, 0, len(s))
	for _, str := range s {
		i, err := strconv.ParseUint(str, base, bitSize)
		if err != nil {
			return nil, err
		}
		r = append(r, Integer(i))
	}
	return r, nil
}

func ParseFloatSlice[Float constraints.Float](s []string, bitSize int) ([]Float, error) {
	r := make([]Float, 0, len(s))
	for _, str := range s {
		f, err := strconv.ParseFloat(str, bitSize)
		if err != nil {
			return nil, err
		}
		r = append(r, Float(f))
	}
	return r, nil
}
