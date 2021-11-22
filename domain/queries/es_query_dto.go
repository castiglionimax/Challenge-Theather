package queries

import (
	"strings"

	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"
)

type EsQuery struct {
	Equals      []FieldValue      `json:"equals"`
	Range_date  *FieldValueRange  `json:"range_date"`
	Range_price *FieldValueRange  `json:"range_price"`
	Orderby     []FieldValueOrder `json:"orderby"`
	Id          FieldValue        `json:"id"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

type FieldValueRange struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

type FieldValueOrder struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

func (q EsQuery) ValidateOrderValue() *errors.RestErr {
	if q.Orderby != nil {
		for _, v := range q.Orderby {
			v.Value = strings.TrimSpace(v.Value)
			if strings.ToUpper(v.Value) != "ASC" {
				if strings.ToUpper(v.Value) != "DESC" {
					return errors.NewBadRequestError("invalid order value")

				}
			}
		}

	}

	return nil
}
