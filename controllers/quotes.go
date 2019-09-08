package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/quote/core"
)

type Quotes struct {
}

func (req *Quotes) Get(ctx context.Requester) (int, interface{}) {
	results := core.GetInvoices(1, 10)

	return http.StatusOK, results
}

// @Title GetQuotes
// @Description Gets the quotes
// @Success 200 {[]core.Entity} []core.Entity
// @router /all/:pagesize [get]
func (req *Quotes) Search(ctx context.Requester) (int, interface{}) {
	page, size := ctx.GetPageData()

	results := core.GetInvoices(page, size)

	return http.StatusOK, results
}

// @Title GetQuote
// @Description Gets the requested Entity
// @Param	key			path	husk.Key 	true		"Key of the entity you require"
// @Success 200 {core.Entity} core.Entity
// @router /:key [get]
func (req *Quotes) View(ctx context.Requester) (int, interface{}) {
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	record, err := core.GetInvoice(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, record
}

// @Title CreateQuote
// @Description Creates a invoice
// @Param	body		body 	logic.Invoice	true		"Invoice object"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *Quotes) Create(ctx context.Requester) (int, interface{}) {
	var entry core.Invoice
	err := ctx.Body(&entry)

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec, err := entry.Create()

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, rec
}

// @Title UpdateQuote
// @Description Updates a Quote
// @Param	body		body 	core.Invoice	true		"body for invoice content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *Quotes) Update(ctx context.Requester) (int, interface{}) {
	body := &core.Invoice{}
	key, err := ctx.GetKeyedRequest(body)

	if err != nil {
		return http.StatusBadRequest, err
	}

	err = body.Update(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}
