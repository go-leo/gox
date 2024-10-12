package internal

import (
	"fmt"
	"reflect"
	"strings"
)

var BoolMap = map[bool]any{
	true:  1,
	false: 0,
}

func AppendEqual(wheres []string, field string, value any) []string {
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.String:
		return append(wheres, fmt.Sprintf("%s = '%s'", field, value))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32, reflect.Float64:
		return append(wheres, fmt.Sprintf("%s = %s", field, value))
	case reflect.Bool:
		return append(wheres, fmt.Sprintf("%s = %d", field, BoolMap[value.(bool)]))
	case reflect.Struct:

	default:
		return append(wheres, fmt.Sprintf("%s = %d", field, value))
	}
}

func AppendLike(wheres []string, field string, value string) []string {
	return append(wheres, fmt.Sprintf("%s LIKE '%s'", field, value))
}

func AppendIn(wheres []string, field string, values []string) []string {
	if len(values) > 0 {
		return append(wheres, fmt.Sprintf("%s IN (%s)", field, strings.Join(values, ",")))
	}
	return wheres
}

func AppendRange(wheres []string, field, minValue, maxValue string) []string {
	if len(minValue) > 0 && len(maxValue) > 0 {
		return append(wheres, fmt.Sprintf("%s >= %s AND %s <= %s", field, minValue, field, maxValue))
	}
	return wheres
}

func AppendIsNotNull(wheres []string, field string) []string {
	return append(wheres, fmt.Sprintf("%s IS NOT NULL", field))
}

func AppendJsonArrayContains(wheres []string, field string, value string) []string {
	return append(wheres, fmt.Sprintf("JSON_ARRAY_CONTAINS(%s, %s)", field, value))
}

func AppendLikes(wheres []string, field string, values []string) []string {
	if len(values) > 0 {
		var likes []string
		for _, item := range values {
			likes = append(likes, fmt.Sprintf("%s LIKE '%s'", field, "%"+strings.Trim(item, "%")+"%"))
		}
		return append(wheres, fmt.Sprintf("(%s)", strings.Join(likes, " OR ")))
	}
	return wheres
}

func And(wheres []string) string {
	if len(wheres) > 0 {
		return strings.Join(wheres, " AND ")
	}
	return ""
}

func Or(wheres []string) string {
	if len(wheres) > 0 {
		return strings.Join(wheres, " OR ")
	}
	return ""
}

func Join(wheres []string, join func(wheres []string) string) string {
	if len(wheres) > 0 {
		return fmt.Sprintf("WHERE %s", join(wheres))
	}
	return ""
}
