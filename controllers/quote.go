package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/quote/core"
)

type QuoteController struct {
	xontrols.APICtrl
}

// @Title GetQuotes
// @Description Gets the quotes
// @Success 200 {[]core.Entity} []core.Entity
// @router /all/:pagesize [get]
func (req *QuoteController) Get() {
	page, size := req.GetPageData()

	results := core.GetInvoices(page, size)

	req.Serve(http.StatusOK, nil, results)
}

// @Title GetQuote
// @Description Gets the requested Entity
// @Param	key			path	husk.Key 	true		"Key of the entity you require"
// @Success 200 {core.Entity} core.Entity
// @router /:key [get]
func (req *QuoteController) GetByID() {
	key, err := husk.ParseKey(req.FindParam("key"))

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	record, err := core.GetInvoice(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, record)
}

// @Title CreateQuote
// @Description Creates a invoice
// @Param	body		body 	logic.Invoice	true		"Invoice object"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *QuoteController) Post() {
	var entry core.Invoice
	err := req.Body(&entry)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec, err := entry.Create()

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec)
}

// @Title UpdateQuote
// @Description Updates a Quote
// @Param	body		body 	core.Invoice	true		"body for invoice content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *QuoteController) Put() {
	body := &core.Invoice{}
	key, err := req.GetKeyedRequest(body)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	err = body.Update(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, nil)
}
