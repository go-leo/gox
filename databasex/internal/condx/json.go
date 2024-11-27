package condx

import (
	"fmt"
	"strings"
)

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
