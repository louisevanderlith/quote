package routers

// @APIVersion 1.0.0
// @Title Entity.API
// @Description API used to access and modify enity details.

import (
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/quote/controllers"

	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(poxy *droxolite.Epoxy) {
	//Quote
	quoteCtrl := &controllers.QuoteController{}
	quoteGroup := droxolite.NewRouteGroup("quote", quoteCtrl)
	quoteGroup.AddRoute("/", "POST", roletype.User, quoteCtrl.Post)
	quoteGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.User, quoteCtrl.GetByID)
	quoteGroup.AddRoute("/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, quoteCtrl.Get)
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
