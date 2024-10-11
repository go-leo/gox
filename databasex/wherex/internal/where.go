package internal

import (
	"fmt"
	"strings"
)

func AppendEqual(wheres []string, field string, value string) []string {
	return append(wheres, fmt.Sprintf("%s = %s", field, value))
}

func AppendLike(wheres []string, field string, value string) []string {
	return append(wheres, fmt.Sprintf("%s LIKE '%s'", field, "%"+value+"%"))
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
