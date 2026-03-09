package builder

import (
	"Cherry/dsl"
	"fmt"
	"strings"
)

func BuildUpdate(obj dsl.Update) string {

	var exp []string

	for col, val := range obj.Set {

		switch v := val.(type) {

		case string:
			exp = append(exp, fmt.Sprintf("%s = '%s'", col, v))

		default:
			exp = append(exp, fmt.Sprintf("%s = %v", col, v))
		}
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s",
		obj.Entity,
		strings.Join(exp, ", "),
	)

	if obj.Where != "" {
		query += " WHERE " + obj.Where
	}

	return query
}
