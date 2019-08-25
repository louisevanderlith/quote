package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/quote/core"
)

type QuoteController struct {
}

// @Title GetQuotes
// @Description Gets the quotes
// @Success 200 {[]core.Entity} []core.Entity
// @router /all/:pagesize [get]
func (req *QuoteController) Get(ctx context.Contexer) (int, interface{}) {
	page, size := ctx.GetPageData()

	results := core.GetInvoices(page, size)

	return http.StatusOK, results
}

// @Title GetQuote
// @Description Gets the requested Entity
// @Param	key			path	husk.Key 	true		"Key of the entity you require"
// @Success 200 {core.Entity} core.Entity
// @router /:key [get]
func (req *QuoteController) GetByID(ctx context.Contexer) (int, interface{}) {
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
func (req *QuoteController) Post(ctx context.Contexer) (int, interface{}) {
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
func (req *QuoteController) Put(ctx context.Contexer) (int, interface{}) {
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
