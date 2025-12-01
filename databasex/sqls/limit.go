package sqls

import "fmt"

func Limit(n int) string {
	return fmt.Sprintf("LIMIT %d", n)
}
