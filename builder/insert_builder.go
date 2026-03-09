package builder

import (
	"Cherry/dsl"
	"fmt"
	"strings"
)

func BuildInsert(obj dsl.Insert) string {

	fields := strings.Join(obj.Fields, ", ")

	var values []string

	for _, v := range obj.Values {

		switch val := v.(type) {

		case string:
			values = append(values, fmt.Sprintf("'%s'", val))

		default:
			values = append(values, fmt.Sprintf("%v", val))
		}
	}

	return fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		obj.Entity,
		fields,
		strings.Join(values, ", "),
	)
}