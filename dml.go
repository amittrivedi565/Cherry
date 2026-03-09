package main

import (
	"fmt"
	"strings"
)

type Insert struct {
	Op     OperationName `json:"op"`
	Entity string        `json:"entity"`
	Fields []string      `json:"fields"`
	Values []interface{} `json:"values"`
}

type Update struct {
	Op     string                 `json:"op"`
	Entity string                 `json:"entity"`
	Set    map[string]interface{} `json:"set"`
	Where  string                 `json:"where"`
}

type Delete struct {
	Op     string `json:"op"`
	Entity string `json:"entity"`
	Where  string `json:"where"`
}

func BuildInsert(obj Insert) string {

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

func BuildUpdate(obj Update) string {

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

func BuildDelete(obj Delete) string {

	query := fmt.Sprintf("DELETE FROM %s", obj.Entity)

	if obj.Where != "" {
		query += " WHERE " + obj.Where
	}

	return query
}
