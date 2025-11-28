package condx

import (
	"fmt"
	"strings"
)

// Combiner 定义了如何组合多个条件
type Combiner interface {
	Combine(conditions []string) string
}

// combinerFunc 是 Combiner 接口的函数实现
type combinerFunc func(conditions []string) string

// Combine 实现 Combiner 接口
func (f combinerFunc) Combine(conditions []string) string {
	return f(conditions)
}

// And 返回一个用 AND 连接条件的组合器
func And() Combiner {
	return combinerFunc(func(conditions []string) string {
		if len(conditions) <= 0 {
			return ""
		}
		return strings.Join(conditions, " AND ")
	})
}

// Or 返回一个用 OR 连接条件的组合器
func Or() Combiner {
	return combinerFunc(func(conditions []string) string {
		if len(conditions) <= 0 {
			return ""
		}
		return strings.Join(conditions, " OR ")
	})
}

// Where 构建 WHERE 子句
func Where(conditions []string, op Combiner) string {
	if len(conditions) <= 0 {
		return ""
	}
	return fmt.Sprintf("WHERE %s", op.Combine(conditions))
}

// Having 构建 HAVING 子句
func Having(conditions []string, op Combiner) string {
	if len(conditions) <= 0 {
		return ""
	}
	return fmt.Sprintf("HAVING %s", op.Combine(conditions))
}

// Eq 添加等于条件 (=)
func Eq(conditions []string, field, value string) []string {
	return append(conditions, fmt.Sprintf("(%s = %s)", field, value))
}

// MustEq 添加必须的等于条件 (=)，当值为空时不添加
func MustEq(conditions []string, field, value string) []string {
	if len(value) <= 0 {
		return conditions
	}
	return Eq(conditions, field, value)
}

// Ne 添加不等于条件 (<>)
func Ne(conditions []string, field, value string) []string {
	return append(conditions, fmt.Sprintf("(%s <> %s)", field, value))
}

// MustNe 添加必须的不等于条件 (<>)，当值为空时不添加
func MustNe(conditions []string, field, value string) []string {
	if len(value) <= 0 {
		return conditions
	}
	return Ne(conditions, field, value)
}

// Gt 添加大于条件 (>)
func Gt(conditions []string, field, value string) []string {
	return append(conditions, fmt.Sprintf("(%s > %s)", field, value))
}

// MustGt 添加必须的大于条件 (>)，当值为空时不添加
func MustGt(conditions []string, field, value string) []string {
	if len(value) <= 0 {
		return conditions
	}
	return Gt(conditions, field, value)
}

// Lt 添加小于条件 (<)
func Lt(conditions []string, field, value string) []string {
	return append(conditions, fmt.Sprintf("(%s < %s)", field, value))
}

// MustLt 添加必须的小于条件 (<)，当值为空时不添加
func MustLt(conditions []string, field, value string) []string {
	if len(value) <= 0 {
		return conditions
	}
	return Lt(conditions, field, value)
}

// Ge 添加大于等于条件 (>=)
func Ge(conditions []string, field, value string) []string {
	return append(conditions, fmt.Sprintf("(%s >= %s)", field, value))
}

// MustGe 添加必须的大于等于条件 (>=)，当值为空时不添加
func MustGe(conditions []string, field, value string) []string {
	if len(value) <= 0 {
		return conditions
	}
	return Ge(conditions, field, value)
}

// Le 添加小于等于条件 (<=)
func Le(conditions []string, field, value string) []string {
	return append(conditions, fmt.Sprintf("(%s <= %s)", field, value))
}

// MustLe 添加必须的小于等于条件 (<=)，当值为空时不添加
func MustLe(conditions []string, field, value string) []string {
	if len(value) <= 0 {
		return conditions
	}
	return Le(conditions, field, value)
}

// Between 添加 BETWEEN 条件
func Between(conditions []string, field, minValue, maxValue string) []string {
	return append(conditions, fmt.Sprintf("(%s BETWEEN %s AND %s)", field, minValue, maxValue))
}

// MustBetween 添加必须的 BETWEEN 条件，当最小值或最大值为空时不添加
func MustBetween(conditions []string, field, minValue, maxValue string) []string {
	if len(minValue) <= 0 || len(maxValue) <= 0 {
		return conditions
	}
	return Between(conditions, field, minValue, maxValue)
}

// NotBetween 添加 NOT BETWEEN 条件
func NotBetween(conditions []string, field, minValue, maxValue string) []string {
	return append(conditions, fmt.Sprintf("(%s NOT BETWEEN %s AND %s)", field, minValue, maxValue))
}

// MustNotBetween 添加必须的 NOT BETWEEN 条件，当最小值或最大值为空时不添加
func MustNotBetween(conditions []string, field, minValue, maxValue string) []string {
	if len(minValue) <= 0 || len(maxValue) <= 0 {
		return conditions
	}
	return NotBetween(conditions, field, minValue, maxValue)
}

// Like 添加 LIKE 条件
func Like(conditions []string, field string, value string) []string {
	return append(conditions, fmt.Sprintf("(%s LIKE '%s')", field, value))
}

// MustLike 添加必须的 LIKE 条件，当值为空时不添加
func MustLike(conditions []string, field string, value string) []string {
	if len(value) <= 0 {
		return conditions
	}
	return Like(conditions, field, value)
}

// NotLike 添加 NOT LIKE 条件
func NotLike(conditions []string, field string, value string) []string {
	return append(conditions, fmt.Sprintf("(%s NOT LIKE '%s')", field, value))
}

// MustNotLike 添加必须的 NOT LIKE 条件，当值为空时不添加
func MustNotLike(conditions []string, field string, value string) []string {
	if len(value) <= 0 {
		return conditions
	}
	return NotLike(conditions, field, value)
}

// In 添加 IN 条件
func In(conditions []string, field string, values []string) []string {
	return append(conditions, fmt.Sprintf("(%s IN (%s))", field, strings.Join(values, ",")))
}

// MustIn 添加必须的 IN 条件，当值为空时不添加
func MustIn(conditions []string, field string, values []string) []string {
	if len(values) <= 0 {
		return conditions
	}
	return In(conditions, field, values)
}

// NotIn 添加 NOT IN 条件
func NotIn(conditions []string, field string, values []string) []string {
	return append(conditions, fmt.Sprintf("(%s NOT IN (%s))", field, strings.Join(values, ",")))
}

// MustNotIn 添加必须的 NOT IN 条件，当值为空时不添加
func MustNotIn(conditions []string, field string, values []string) []string {
	if len(values) <= 0 {
		return conditions
	}
	return NotIn(conditions, field, values)
}

// IsNull 添加 IS NULL 条件
func IsNull(conditions []string, field string) []string {
	return append(conditions, fmt.Sprintf("(%s IS NULL)", field))
}

// IsNotNull 添加 IS NOT NULL 条件
func IsNotNull(conditions []string, field string) []string {
	return append(conditions, fmt.Sprintf("(%s IS NOT NULL)", field))
}

// Exists 添加 EXISTS 条件
func Exists(conditions []string, subquery string) []string {
	return append(conditions, fmt.Sprintf("(EXISTS (%s))", subquery))
}

// NotExists 添加 NOT EXISTS 条件
func NotExists(conditions []string, subquery string) []string {
	return append(conditions, fmt.Sprintf("(NOT EXISTS (%s))", subquery))
}

// EqAll 添加 = ALL 条件
func EqAll(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s = ALL (%s))", field, subquery))
}

// EqAny 添加 = ANY 条件
func EqAny(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s = ANY (%s))", field, subquery))
}

// NeAll 添加 <> ALL 条件
func NeAll(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s <> ALL (%s))", field, subquery))
}

// NeAny 添加 <> ANY 条件
func NeAny(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s <> ANY (%s))", field, subquery))
}

// GtAll 添加 > ALL 条件
func GtAll(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s > ALL (%s))", field, subquery))
}

// GtAny 添加 > ANY 条件
func GtAny(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s > ANY (%s))", field, subquery))
}

// LtAll 添加 < ALL 条件
func LtAll(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s < ALL (%s))", field, subquery))
}

// LtAny 添加 < ANY 条件
func LtAny(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s < ANY (%s))", field, subquery))
}

// GeAll 添加 >= ALL 条件
func GeAll(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s >= ALL (%s))", field, subquery))
}

// GeAny 添加 >= ANY 条件
func GeAny(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s >= ANY (%s))", field, subquery))
}

// LeAll 添加 <= ALL 条件
func LeAll(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s <= ALL (%s))", field, subquery))
}

// LeAny 添加 <= ANY 条件
func LeAny(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s <= ANY (%s))", field, subquery))
}
