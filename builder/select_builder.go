package builder

import (
	"Cherry/dsl"
	"fmt"
	"strings"
)

func BuildSelect(object dsl.Select) string {

	fields := "*"
	if len(object.Select) > 0 {
		fields = strings.Join(object.Select, ", ")
	}

	query := fmt.Sprintf("SELECT %s FROM %s %s", fields, object.Entity, object.Alias)

	for _, j := range object.Joins {

		joinType := "JOIN"
		if j.Type != "" {
			joinType = strings.ToUpper(j.Type) + " JOIN"
		}

		query += fmt.Sprintf(
			" %s %s %s ON %s = %s",
			joinType,
			j.Entity,
			j.Alias,
			j.On.Left,
			j.On.Right,
		)
	}

	if object.Where.Condition != nil || len(object.Where.Conditions) > 0 {
		query += " WHERE " + BuildWhere(object.Where)
	}

	return query
}

func BuildWhere(object dsl.WhereNode) string {

	if object.Condition != nil {
		c := object.Condition

		switch v := c.Value.(type) {
		case string:
			return fmt.Sprintf("%s %s '%s'", c.Field, c.Op, v)
		default:
			return fmt.Sprintf("%s %s %v", c.Field, c.Op, v)
		}
	}

	var parts []string

	for _, child := range object.Conditions {
		parts = append(parts, BuildWhere(child))
	}

	if object.Operator == "" {
		return strings.Join(parts, " ")
	}

	return "(" + strings.Join(parts, " "+object.Operator+" ") + ")"
}
