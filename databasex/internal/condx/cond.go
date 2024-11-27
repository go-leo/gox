package condx

import (
	"fmt"
)

func Where(conds []string, op Combiner) string {
	if len(conds) == 0 {
		return ""
	}
	return fmt.Sprintf("WHERE %s", op.Combine(conds))
}

func Having(conds []string, op Combiner) string {
	if len(conds) == 0 {
		return ""
	}
	return fmt.Sprintf("HAVING %s", op.Combine(conds))
}

func On(conds []string) string {
	if len(conds) == 0 {
		return ""
	}
	return fmt.Sprintf("HAVING %s", And().Combine(conds))
}
