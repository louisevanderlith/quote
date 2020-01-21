package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	Invoices husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Invoices: husk.NewTable(Invoice{}, serials.GobSerial{}),
	}
}

func Shutdown() {
	ctx.Invoices.Save()
}
