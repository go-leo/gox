package condx

import (
	"fmt"
	"reflect"
	"strings"
)

var BoolMap = map[bool]any{
	true:  1,
	false: 0,
}

func AppendEqual(conds []string, field string, value any) []string {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.String:
		return append(conds, fmt.Sprintf("%s = '%s'", field, value))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32, reflect.Float64:
		return append(conds, fmt.Sprintf("%s = %s", field, value))
	case reflect.Bool:
		return append(conds, fmt.Sprintf("%s = %d", field, BoolMap[value.(bool)]))
	case reflect.Struct:

	default:
		return append(conds, fmt.Sprintf("%s = %d", field, value))
	}
	return conds
}

func AppendLike(conds []string, field string, value string) []string {
	return append(conds, fmt.Sprintf("%s LIKE '%s'", field, value))
}

func AppendIn(conds []string, field string, values []string) []string {
	if len(values) > 0 {
		return append(conds, fmt.Sprintf("%s IN (%s)", field, strings.Join(values, ",")))
	}
	return conds
}

func AppendRange(conds []string, field, minValue, maxValue string) []string {
	if len(minValue) > 0 && len(maxValue) > 0 {
		return append(conds, fmt.Sprintf("%s >= %s AND %s <= %s", field, minValue, field, maxValue))
	}
	return conds
}

func AppendIsNotNull(conds []string, field string) []string {
	return append(conds, fmt.Sprintf("%s IS NOT NULL", field))
}

func AppendJsonArrayContains(conds []string, field string, value string) []string {
	return append(conds, fmt.Sprintf("JSON_ARRAY_CONTAINS(%s, %s)", field, value))
}

func AppendLikes(conds []string, field string, values []string) []string {
	if len(values) > 0 {
		var likes []string
		for _, item := range values {
			likes = append(likes, fmt.Sprintf("%s LIKE '%s'", field, "%"+strings.Trim(item, "%")+"%"))
		}
		return append(conds, fmt.Sprintf("(%s)", strings.Join(likes, " OR ")))
	}
	return conds
}

func And(conds []string) string {
	if len(conds) > 0 {
		return strings.Join(conds, " AND ")
	}
	return ""
}

func Or(conds []string) string {
	if len(conds) > 0 {
		return strings.Join(conds, " OR ")
	}
	return ""
}

func Where(conds []string, join func(conds []string) string) string {
	if len(conds) > 0 {
		return fmt.Sprintf("WHERE %s", join(conds))
	}
	return ""
}

func Having(conds []string, join func(conds []string) string) string {
	if len(conds) > 0 {
		return fmt.Sprintf("HAVING %s", join(conds))
	}
	return ""
}
