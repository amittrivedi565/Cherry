package dsl

type Delete struct {
	Op     string `json:"op"`
	Entity string `json:"entity"`
	Where  string `json:"where"`
}