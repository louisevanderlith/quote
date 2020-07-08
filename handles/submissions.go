package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/quote/core"
)

func GetSubmissions(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	results, err := core.GetInvoices(1, 10)

	if err != nil {
		log.Println("Get Submissions", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title GetQuotes
// @Description Gets the quotes
// @Success 200 {[]core.Entity} []core.Entity
// @router /all/:pagesize [get]
func SearchSubmissions(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	page, size := ctx.GetPageData()

	results, err := core.GetInvoices(page, size)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title GetQuote
// @Description Gets the requested Entity
// @Param	key			path	husk.Key 	true		"Key of the entity you require"
// @Success 200 {core.Entity} core.Entity
// @router /:key [get]
func ViewSubmission(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println("ParseKey Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	record, err := core.GetInvoice(key)

	if err != nil {
		log.Println("GetInvoice Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(record))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title CreateQuote
// @Description Creates a invoice
// @Param	body		body 	logic.Invoice	true		"Invoice object"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func CreateQuote(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	var entry core.Submission
	err := ctx.Body(&entry)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := entry.Create()

	if err != nil {
		log.Println("Create Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(rec))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title UpdateQuote
// @Description Updates a Quote
// @Param	body		body 	core.Invoice	true		"body for invoice content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func UpdateSubmission(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println("Parse Key Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body := core.Submission{}
	err = ctx.Body(&body)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = body.Update(key)

	if err != nil {
		log.Println("Update Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON("Saved"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
