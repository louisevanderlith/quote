package routers

// @APIVersion 1.0.0
// @Title Entity.API
// @Description API used to access and modify enity details.

import (
	"github.com/louisevanderlith/quote/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
)

func Setup(poxy resins.Epoxi) {
	//Quote
	quoteCtrl := &controllers.QuoteController{}
	quoteGroup := routing.NewRouteGroup("quote", mix.JSON)
	quoteGroup.AddRoute("Create Quote", "", "POST", roletype.User, quoteCtrl.Post)
	quoteGroup.AddRoute("Quote by Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.User, quoteCtrl.GetByID)
	quoteGroup.AddRoute("All Quotes", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, quoteCtrl.Get)
	poxy.AddGroup(quoteGroup)
	/*ctrlmap := EnableFilters(s, host)
	quoteCtrl := controllers.NewQuoteCtrl(ctrlmap)

	beego.Router("/v1/quote", quoteCtrl, "post:Post")
	beego.Router("/v1/quote/:key", quoteCtrl, "get:GetByID")
	beego.Router("/v1/quote/all/:pagesize", quoteCtrl, "get:Get")*/
}

/*
func EnableFilters(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.User
	emptyMap["GET"] = roletype.User

	ctrlmap.Add("/v1/quote", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}*/
