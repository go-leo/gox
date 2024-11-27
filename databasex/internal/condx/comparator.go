package condx

import (
	"fmt"
)

type Comparator interface {
	Compare(conds []string, field string, value any) []string
}

// Equal generates an 'expression1 = expression2' condition
// and appends it to an existing list of conditions.
func Equal(conds []string, field string, value any) []string {
	return append(conds, fmt.Sprintf("%s = %d", field, value))
}

// NotEqual generates an 'expression1 <> expression2' condition
// and appends it to an existing list of conditions.
func NotEqual(conds []string, field string, value any) []string {
	return append(conds, fmt.Sprintf("%s <> %d", field, value))
}

// GreaterThan generates an 'expression1 > expression2' condition
// and appends it to an existing list of conditions.
func GreaterThan(conds []string, field string, value any) []string {
	return append(conds, fmt.Sprintf("%s > %d", field, value))
}

// GreaterThanEqual generates an 'expression1 >= expression2' condition
// and appends it to an existing list of conditions.
func GreaterThanEqual(conds []string, field string, value any) []string {
	return append(conds, fmt.Sprintf("%s >= %d", field, value))
}

// LessThan generates an 'expression1 < expression2' condition
// and appends it to an existing list of conditions.
func LessThan(conds []string, field string, value any) []string {
	return append(conds, fmt.Sprintf("%s < %d", field, value))
}

// LessThanEqual generates an 'expression1 <= expression2' condition
// and appends it to an existing list of conditions.
func LessThanEqual(conds []string, field string, value any) []string {
	return append(conds, fmt.Sprintf("%s <= %d", field, value))
}
