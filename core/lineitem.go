package core

import "github.com/louisevanderlith/husk"

type LineItem struct {
	Description string
	UnitCost    int //we don't use decimals @ the back
	Quantity    int
	SubItems    []string
}

func (l LineItem) Valid() error {
	return husk.ValidateStruct(&l)
}
