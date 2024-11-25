package condx

import (
	"fmt"
)

func Where(conds []string, op Operator) string {
	if len(conds) > 0 {
		return fmt.Sprintf("WHERE %s", op.Apply(conds))
	}
	return ""
}

func Having(conds []string, op Operator) string {
	if len(conds) > 0 {
		return fmt.Sprintf("HAVING %s", op.Apply(conds))
	}
	return ""
}
