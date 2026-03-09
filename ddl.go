package main

import (
	"fmt"
	"strings"
)

type Create struct {
	Op     OperationName `json:"op"`
	Entity string        `json:"entity"`
	Fields []Field       `json:"fields"`
}

type Field struct {
	Name      string `json:"field_name"`
	FieldType string `json:"field_type"`
	Key       string `json:"field_key"`
}

type Drop struct {
	Op     string `json:"op"`
	Entity string `json:"entity"`
}

func BuildCreate(obj Create) string {

	var fields []string

	for _, value := range obj.Fields {
		field := fmt.Sprintf("%s %s %s", value.Name, value.FieldType, value.Key)
		fields = append(fields, field)
	}

	return fmt.Sprintf(
		"CREATE TABLE %s (%s)",
		obj.Entity,
		strings.Join(fields, ", "),
	)
}

func BuildDrop(obj Drop) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS %s", obj.Entity)
}
