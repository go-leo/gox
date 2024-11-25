package condx

import "strings"

type Operator interface {
	Apply(conds []string) string
}

type OperatorFunc func(conds []string) string

func (f OperatorFunc) Apply(conds []string) string {
	return f(conds)
}

func And() Operator {
	return OperatorFunc(func(conds []string) string {
		if len(conds) > 0 {
			return strings.Join(conds, " AND ")
		}
		return ""
	})
}

func Or() Operator {
	return OperatorFunc(func(conds []string) string {
		if len(conds) > 0 {
			return strings.Join(conds, " OR ")
		}
		return ""
	})
}
