package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/husk/validation"
	"time"

	"github.com/louisevanderlith/quote/core/statustype"
)

type Submission struct {
	Status    statustype.Enum
	IsInvoice bool
	DueDate   time.Time
	Number    int
	LineItems []LineItem
}

func (i Submission) Valid() error {
	return validation.Struct(i)
}

func GetInvoices(page, pagesize int) (records.Page, error) {
	return ctx.Submissions.Find(page, pagesize, op.Everything())
}

func GetInvoice(key hsk.Key) (hsk.Record, error) {
	return ctx.Submissions.FindByKey(key)
}

func (i Submission) Create() (hsk.Key, error) {
	return ctx.Submissions.Create(i)
}

func (i Submission) Update(key hsk.Key) error {
	return ctx.Submissions.Update(key, i)
}
