package sqls

import "fmt"

func From(table string) string {
	return fmt.Sprintf("FROM %s", table)
}
