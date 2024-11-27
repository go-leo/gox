package condx

import (
	"fmt"
	"strings"
)

func IsNull(conds []string, field string) []string {
	return append(conds, fmt.Sprintf("%s IS NULL", field))
}

func IsNotNull(conds []string, field string) []string {
	return append(conds, fmt.Sprintf("%s IS NOT NULL", field))
}

func Between(conds []string, field, minValue, maxValue string) []string {
	if len(minValue) > 0 && len(maxValue) > 0 {
		return append(conds, fmt.Sprintf("%s BETWEEN %s AND %s", field, minValue, maxValue))
	}
	return conds
}

func NotBetween(conds []string, field, minValue, maxValue string) []string {
	if len(minValue) > 0 && len(maxValue) > 0 {
		return append(conds, fmt.Sprintf("%s NOT BETWEEN %s AND %s", field, minValue, maxValue))
	}
	return conds
}

func In(conds []string, field string, values []string) []string {
	if len(values) > 0 {
		return append(conds, fmt.Sprintf("%s IN (%s)", field, strings.Join(values, ",")))
	}
	return conds
}

func NotIn(conds []string, field string, values []string) []string {
	if len(values) > 0 {
		return append(conds, fmt.Sprintf("%s NOT IN (%s)", field, strings.Join(values, ",")))
	}
	return conds
}

func Like(conds []string, field string, value string) []string {
	return append(conds, fmt.Sprintf("%s LIKE '%s'", field, value))
}

func NotLike(conds []string, field string, value string) []string {
	return append(conds, fmt.Sprintf("%s NOT LIKE '%s'", field, value))
}

// All generates an 'ALL (subquery)' condition
func All(subquery string) string {
	if len(subquery) <= 0 {
		return ""
	}
	return fmt.Sprintf("ALL (%s)", subquery)
}

// Any generates an 'ANY (subquery)' condition
func Any(subquery string) string {
	if len(subquery) <= 0 {
		return ""
	}
	return fmt.Sprintf("ANY (%s)", subquery)
}

// Exists generates an 'EXISTS (subquery)' condition
func Exists(conds []string, subquery string) []string {
	return append(conds, fmt.Sprintf("EXISTS (%s)", subquery))
}

// NotExists generates an 'NOT EXISTS (subquery)' condition
func NotExists(conds []string, subquery string) []string {
	return append(conds, fmt.Sprintf("EXISTS (%s)", subquery))
}
