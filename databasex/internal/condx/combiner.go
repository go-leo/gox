package condx

import (
	"strings"
)

type Combiner interface {
	Combine(conds []string) string
}

type combinerFunc func(conds []string) string

func (f combinerFunc) Combine(conds []string) string {
	return f(conds)
}

func And() Combiner {
	return combinerFunc(func(conds []string) string {
		if len(conds) > 0 {
			return strings.Join(conds, " AND ")
		}
		return ""
	})
}

func Or() Combiner {
	return combinerFunc(func(conds []string) string {
		if len(conds) > 0 {
			return strings.Join(conds, " OR ")
		}
		return ""
	})
}
