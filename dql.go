package main

import (
	"fmt"
	"strings"
)

type Select struct {
	Op     OperationName `json:"op"`
	Entity string        `json:"entity"`
	Alias  string        `json:"alias"`
	Select []string      `json:"select"`
	Joins  []Join        `json:"joins"`
	Where  WhereNode     `json:"where"`
}

type Join struct {
	Type   string `json:"type"`
	Entity string `json:"entity"`
	Alias  string `json:"alias"`
	On     On     `json:"on"`
}

type On struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

type Condition struct {
	Field string      `json:"field"`
	Op    string      `json:"op"`
	Value interface{} `json:"value"`
}

type WhereNode struct {
	Operator   string      `json:"operator,omitempty"`
	Conditions []WhereNode `json:"conditions,omitempty"`
	Condition  *Condition  `json:"condition,omitempty"`
}

func BuildSelect(object Select) string {

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

func BuildWhere(object WhereNode) string {

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
