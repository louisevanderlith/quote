package core

import "github.com/louisevanderlith/husk"

type context struct {
	Invoices husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Invoices: husk.NewTable(new(Invoice)),
	}
}

func Shutdown() {
	ctx.Invoices.Save()
}
