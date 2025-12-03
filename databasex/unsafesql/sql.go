package unsafesql

import (
	"errors"
	"fmt"
	"strings"
)

type Combiner interface {
	Combine(conditions []string) string
}

type combinerFunc func(conditions []string) string

func (f combinerFunc) Combine(conditions []string) string {
	return f(conditions)
}

func And() Combiner {
	return combinerFunc(func(conditions []string) string {
		if len(conditions) <= 0 {
			return ""
		}
		return strings.Join(conditions, " AND ")
	})
}

func Or() Combiner {
	return combinerFunc(func(conditions []string) string {
		if len(conditions) <= 0 {
			return ""
		}
		return strings.Join(conditions, " OR ")
	})
}

func Where(conditions []string, op Combiner) string {
	if len(conditions) <= 0 {
		return ""
	}
	return fmt.Sprintf("WHERE %s", op.Combine(conditions))
}

func MustWhere(conditions []string, op Combiner) string {
	if len(conditions) <= 0 {
		panic(errors.New("sqls: conditions is empty"))
	}
	return Where(conditions, op)
}

func Having(conditions []string, op Combiner) string {
	if len(conditions) <= 0 {
		return ""
	}
	return fmt.Sprintf("HAVING %s", op.Combine(conditions))
}

func MustHaving(conditions []string, op Combiner) string {
	if len(conditions) <= 0 {
		panic(errors.New("sqls: conditions is empty"))
	}
	return Having(conditions, op)
}

func Eq(conditions []string, field, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		return conditions
	}
	return append(conditions, fmt.Sprintf("(%s = %s)", field, value))
}

func MustEq(conditions []string, field, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		panic(errors.New("sqls: field or value is empty"))
	}
	return Eq(conditions, field, value)
}

func Ne(conditions []string, field, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		return conditions
	}
	return append(conditions, fmt.Sprintf("(%s <> %s)", field, value))
}

func MustNe(conditions []string, field, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		panic(errors.New("sqls: field or value is empty"))
	}
	return Ne(conditions, field, value)
}

func Gt(conditions []string, field, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		return conditions
	}
	return append(conditions, fmt.Sprintf("(%s > %s)", field, value))
}

func MustGt(conditions []string, field, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		panic(errors.New("sqls: field or value is empty"))
	}
	return Gt(conditions, field, value)
}

func Lt(conditions []string, field, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		return conditions
	}
	return append(conditions, fmt.Sprintf("(%s < %s)", field, value))
}

func MustLt(conditions []string, field, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		panic(errors.New("sqls: field or value is empty"))
	}
	return Lt(conditions, field, value)
}

func Ge(conditions []string, field, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		return conditions
	}
	return append(conditions, fmt.Sprintf("(%s >= %s)", field, value))
}

func MustGe(conditions []string, field, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		panic(errors.New("sqls: field or value is empty"))
	}
	return Ge(conditions, field, value)
}

func Le(conditions []string, field, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		return conditions
	}
	return append(conditions, fmt.Sprintf("(%s <= %s)", field, value))
}

func MustLe(conditions []string, field, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		panic(errors.New("sqls: field or value is empty"))
	}
	return Le(conditions, field, value)
}

func Between(conditions []string, field, minValue, maxValue string) []string {
	if len(field) <= 0 || len(minValue) <= 0 || len(maxValue) <= 0 {
		return conditions
	}
	return append(conditions, fmt.Sprintf("(%s BETWEEN %s AND %s)", field, minValue, maxValue))
}

func MustBetween(conditions []string, field, minValue, maxValue string) []string {
	if len(field) <= 0 || len(minValue) <= 0 || len(maxValue) <= 0 {
		panic(errors.New("sqls: field or min or max value is empty"))
	}
	return Between(conditions, field, minValue, maxValue)
}

func NotBetween(conditions []string, field, minValue, maxValue string) []string {
	if len(field) <= 0 || len(minValue) <= 0 || len(maxValue) <= 0 {
		return conditions
	}
	return append(conditions, fmt.Sprintf("(%s NOT BETWEEN %s AND %s)", field, minValue, maxValue))
}

func MustNotBetween(conditions []string, field, minValue, maxValue string) []string {
	if len(field) <= 0 || len(minValue) <= 0 || len(maxValue) <= 0 {
		panic(errors.New("sqls: field or min or max value is empty"))
	}
	return NotBetween(conditions, field, minValue, maxValue)
}

func Like(conditions []string, field string, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		return conditions
	}
	return append(conditions, fmt.Sprintf("(%s LIKE '%s')", field, value))
}

func MustLike(conditions []string, field string, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		panic(errors.New("sqls: field or value is empty"))
	}
	return Like(conditions, field, value)
}

func NotLike(conditions []string, field string, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		return conditions
	}
	return append(conditions, fmt.Sprintf("(%s NOT LIKE '%s')", field, value))
}

func MustNotLike(conditions []string, field string, value string) []string {
	if len(field) <= 0 || len(value) <= 0 {
		panic(errors.New("sqls: field or value is empty"))
	}
	return NotLike(conditions, field, value)
}

func In(conditions []string, field string, values []string) []string {
	if len(field) <= 0 || len(values) <= 0 {
		return conditions
	}
	return append(conditions, fmt.Sprintf("(%s IN (%s))", field, strings.Join(values, ", ")))
}

func MustIn(conditions []string, field string, values []string) []string {
	if len(field) <= 0 || len(values) <= 0 {
		panic(errors.New("sqls: field or values is empty"))
	}
	return In(conditions, field, values)
}

func NotIn(conditions []string, field string, values []string) []string {
	if len(field) <= 0 || len(values) <= 0 {
		return conditions
	}
	return append(conditions, fmt.Sprintf("(%s NOT IN (%s))", field, strings.Join(values, ", ")))
}

func MustNotIn(conditions []string, field string, values []string) []string {
	if len(field) <= 0 || len(values) <= 0 {
		panic(errors.New("sqls: field or values is empty"))
	}
	return NotIn(conditions, field, values)
}

func IsNull(conditions []string, field string) []string {
	return append(conditions, fmt.Sprintf("(%s IS NULL)", field))
}

func MustIsNull(conditions []string, field string) []string {
	if len(field) <= 0 {
		panic(errors.New("sqls: field is empty"))
	}
	return IsNull(conditions, field)
}

func IsNotNull(conditions []string, field string) []string {
	return append(conditions, fmt.Sprintf("(%s IS NOT NULL)", field))
}

func MustIsNotNull(conditions []string, field string) []string {
	if len(field) <= 0 {
		panic(errors.New("sqls: field is empty"))
	}
	return IsNotNull(conditions, field)
}

func Exists(conditions []string, subquery string) []string {
	return append(conditions, fmt.Sprintf("(EXISTS (%s))", subquery))
}

func MustExists(conditions []string, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return Exists(conditions, subquery)
}

func NotExists(conditions []string, subquery string) []string {
	return append(conditions, fmt.Sprintf("(NOT EXISTS (%s))", subquery))
}

func MustNotExists(conditions []string, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return NotExists(conditions, subquery)
}

func EqAll(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s = ALL (%s))", field, subquery))
}

func MustEqAll(conditions []string, field, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return EqAll(conditions, field, subquery)
}

func EqAny(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s = ANY (%s))", field, subquery))
}

func MustEqAny(conditions []string, field, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return EqAny(conditions, field, subquery)
}

func NeAll(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s <> ALL (%s))", field, subquery))
}

func MustNeAll(conditions []string, field, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return NeAll(conditions, field, subquery)
}

func NeAny(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s <> ANY (%s))", field, subquery))
}

func MustNeAny(conditions []string, field, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return NeAny(conditions, field, subquery)
}

func GtAll(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s > ALL (%s))", field, subquery))
}

func MustGtAll(conditions []string, field, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return GtAll(conditions, field, subquery)
}

func GtAny(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s > ANY (%s))", field, subquery))
}

func MustGtAny(conditions []string, field, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return GtAny(conditions, field, subquery)
}

func LtAll(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s < ALL (%s))", field, subquery))
}

func MustLtAll(conditions []string, field, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return LtAll(conditions, field, subquery)
}

func LtAny(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s < ANY (%s))", field, subquery))
}

func MustLtAny(conditions []string, field, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return LtAny(conditions, field, subquery)
}

func GeAll(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s >= ALL (%s))", field, subquery))
}

func MustGeAll(conditions []string, field, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return GeAll(conditions, field, subquery)
}

func GeAny(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s >= ANY (%s))", field, subquery))
}

func MustGeAny(conditions []string, field, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return GeAny(conditions, field, subquery)
}

func LeAll(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s <= ALL (%s))", field, subquery))
}

func MustLeAll(conditions []string, field, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return LeAll(conditions, field, subquery)
}

func LeAny(conditions []string, field, subquery string) []string {
	return append(conditions, fmt.Sprintf("(%s <= ANY (%s))", field, subquery))
}

func MustLeAny(conditions []string, field, subquery string) []string {
	if len(subquery) <= 0 {
		panic(errors.New("sqls: subquery is empty"))
	}
	return LeAny(conditions, field, subquery)
}

func Select(fields ...string) string {
	if len(fields) <= 0 {
		return "SELECT *"
	}
	return fmt.Sprintf("SELECT %s", strings.Join(fields, ", "))
}

func From(table string) string {
	if len(table) == 0 {
		panic(errors.New("sqls: table is empty"))
	}
	return fmt.Sprintf("FROM %s", table)
}

func LeftJoin(table string, conditions ...string) string {
	if len(table) <= 0 || len(conditions) <= 0 {
		panic(errors.New("sqls: conditions is empty"))
	}
	return fmt.Sprintf("LEFT JOIN %s ON %s", table, strings.Join(conditions, " AND "))
}

func RightJoin(table string, conditions ...string) string {
	if len(table) <= 0 || len(conditions) <= 0 {
		panic(errors.New("sqls: conditions is empty"))
	}
	return fmt.Sprintf("RIGHT JOIN %s ON %s", table, strings.Join(conditions, " AND "))
}

func GroupBy(fields []string) string {
	if len(fields) <= 0 {
		return ""
	}
	return fmt.Sprintf("GROUP BY %s", strings.Join(fields, ", "))
}

func MustGroupBy(fields []string) string {
	if len(fields) <= 0 {
		panic(errors.New("sqls: fields is empty"))
	}
	return GroupBy(fields)
}

func Limit(n int) string {
	if n < 0 {
		return ""
	}
	return fmt.Sprintf("LIMIT %d", n)
}

func MustLimit(n int) string {
	if n < 0 {
		panic(errors.New("sqls: n is less than 0"))
	}
	return Limit(n)
}

func Offset(n int) string {
	if n < 0 {
		return ""
	}
	return fmt.Sprintf("OFFSET %d", n)
}

func MustOffset(n int) string {
	if n < 0 {
		panic(errors.New("sqls: n is less than 0"))
	}
	return Offset(n)
}

func OrderBy(fields []string) string {
	if len(fields) == 0 {
		return ""
	}
	return fmt.Sprintf("ORDER BY %s", strings.Join(fields, ", "))
}

func MustOrderBy(fields []string) string {
	if len(fields) <= 0 {
		panic(errors.New("sqls: fields is empty"))
	}
	return OrderBy(fields)
}
