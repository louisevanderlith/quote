package routers

// @APIVersion 1.0.0
// @Title Entity.API
// @Description API used to access and modify enity details.

import (
	"github.com/louisevanderlith/quote/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e resins.Epoxi) {
	e.JoinBundle("/", roletype.User, mix.JSON, &controllers.Quotes{})
	//Quote
	/*
		quoteCtrl := &controllers.QuoteController{}
		quoteGroup := routing.NewRouteGroup("quote", mix.JSON)
		quoteGroup.AddRoute("Create Quote", "", "POST", roletype.User, quoteCtrl.Post)
		quoteGroup.AddRoute("Quote by Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.User, quoteCtrl.GetByID)
		quoteGroup.AddRoute("All Quotes", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, quoteCtrl.Get)
		e.AddBundle(quoteGroup)*/
}
