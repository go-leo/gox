package sqls

import (
	"fmt"
	"strings"
)

func GroupBy(fields []string) string {
	return fmt.Sprintf("GROUP BY %s", strings.Join(fields, ","))
}

func MustGroupBy(fields []string) string {
	if len(fields) == 0 {
		return ""
	}
	return GroupBy(fields)
}
