package core

import (
	"time"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/quote/core/statustype"
)

type Submission struct {
	Status    statustype.Enum
	IsInvoice  bool
	DueDate   time.Time
	Number    int
	LineItems []LineItem
}

func (i Submission) Valid() error {
	return husk.ValidateStruct(i)
}

func GetInvoices(page, pagesize int) (husk.Collection, error) {
	return ctx.Submissions.Find(page, pagesize, husk.Everything())
}

func GetInvoice(key husk.Key) (husk.Recorder, error) {
	return ctx.Submissions.FindByKey(key)
}

func (i Submission) Create() (husk.Recorder, error) {
	rec := ctx.Submissions.Create(i)

	if rec.Error != nil {
		return nil, rec.Error
	}

	err := ctx.Submissions.Save()

	if err != nil {
		return nil, err
	}

	return rec.Record, nil
}

func (i Submission) Update(key husk.Key) error {
	entry, err := ctx.Submissions.FindByKey(key)

	if err != nil {
		return err
	}

	err = entry.Set(i)

	if err != nil {
		return err
	}

	err = ctx.Submissions.Update(entry)

	if err != nil {
		return err
	}

	return ctx.Submissions.Save()
}