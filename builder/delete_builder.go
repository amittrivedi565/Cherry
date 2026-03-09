package builder

import (
	"Cherry/dsl"
	"fmt"
)


func BuildDelete(obj dsl.Delete) string {

	query := fmt.Sprintf("DELETE FROM %s", obj.Entity)

	if obj.Where != "" {
		query += " WHERE " + obj.Where
	}

	return query
}
