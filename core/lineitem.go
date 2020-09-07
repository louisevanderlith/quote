package core

import (
	"github.com/louisevanderlith/husk/validation"
)

type LineItem struct {
	Description string
	UnitCost    int //we don't use decimals @ the back
	Quantity    int
	SubItems    []string
}

func (l LineItem) Valid() error {
	return validation.Struct(l)
}
