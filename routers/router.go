package routers

// @APIVersion 1.0.0
// @Title Entity.API
// @Description API used to access and modify enity details.

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/quote/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilters(s)
	quoteCtrl := controllers.NewQuoteCtrl(ctrlmap)

	beego.Router("/v1/quote", quoteCtrl, "post:Post")
	beego.Router("/v1/quote/:key", quoteCtrl, "get:GetByID")
	beego.Router("/v1/quote/all/:pagesize", quoteCtrl, "get:Get")
}

func EnableFilters(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.User
	emptyMap["GET"] = roletype.User

	ctrlmap.Add("/v1/quote", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	return ctrlmap
}
