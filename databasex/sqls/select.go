package sqls

import (
	"fmt"
	"strings"
)

func Select(fields []string) string {
	return fmt.Sprintf("SELECT %s", strings.Join(fields, ","))
}

func MustSelect(fields []string) string {
	if len(fields) == 0 {
		return ""
	}
	return Select(fields)
}
