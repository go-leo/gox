package sqls

import (
	"fmt"
	"strings"
)

func OrderBy(fields []string) string {
	return fmt.Sprintf("ORDER BY %s", strings.Join(fields, ","))
}

func MustOrderBy(fields []string) string {
	if len(fields) == 0 {
		return ""
	}
	return OrderBy(fields)
}
