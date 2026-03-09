package dsl

type Select struct {
	Op     string    `json:"op"`
	Entity string    `json:"entity"`
	Alias  string    `json:"alias"`
	Select []string  `json:"select"`
	Joins  []Join    `json:"joins"`
	Where  WhereNode `json:"where"`
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
