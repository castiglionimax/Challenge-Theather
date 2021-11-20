package queries

type EsQuery struct {
	Equals []FieldValue `json:"equals"`
	Range  []FieldValue `json:"range"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
