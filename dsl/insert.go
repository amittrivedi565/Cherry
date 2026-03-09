package dsl

type Insert struct {
	Op     string        `json:"op"`
	Entity string        `json:"entity"`
	Fields []string      `json:"fields"`
	Values []interface{} `json:"values"`
}
