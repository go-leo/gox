package sqls

import "fmt"

func Offset(n int) string {
	return fmt.Sprintf("OFFSET %d", n)
}
