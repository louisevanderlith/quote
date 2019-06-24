package core

import (
	"time"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/quote/core/statustype"
)

type Invoice struct {
	Status    statustype.Enum
	IsMature  bool
	DueDate   time.Time
	Number    int
	LineItems []LineItem
}

func (i Invoice) Valid() (bool, error) {
	return husk.ValidateStruct(&i)
}

func GetInvoices(page, pagesize int) husk.Collection {
	return ctx.Invoices.Find(page, pagesize, husk.Everything())
}

func GetInvoice(key husk.Key) (husk.Recorder, error) {
	return ctx.Invoices.FindByKey(key)
}

func (i Invoice) Create() (husk.Recorder, error) {
	rec := ctx.Invoices.Create(i)

	if rec.Error != nil {
		return nil, rec.Error
	}

	ctx.Invoices.Save()

	return rec.Record, nil
}

func (i Invoice) Update(key husk.Key) error {
	entry, err := ctx.Invoices.FindByKey(key)

	if err != nil {
		return err
	}

	err = entry.Set(i)

	if err != nil {
		return err
	}

	defer ctx.Invoices.Save()
	return ctx.Invoices.Update(entry)
}
