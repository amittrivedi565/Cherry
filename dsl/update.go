package dsl

type Update struct {
	Op     string                 `json:"op"`
	Entity string                 `json:"entity"`
	Set    map[string]interface{} `json:"set"`
	Where  string                 `json:"where"`
}
