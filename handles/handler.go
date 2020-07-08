package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, secureUrl string) http.Handler {
	/*
		//Quote
				quoteCtrl := &controllers.QuoteController{}
				quoteGroup := routing.NewRouteGroup("quote", mix.JSON)
				quoteGroup.AddRoute("Create Quote", "", "POST", roletype.User, quoteCtrl.Post)
				quoteGroup.AddRoute("Quote by Key", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.User, quoteCtrl.GetByID)
				quoteGroup.AddRoute("All Quotes", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, quoteCtrl.Get)
				e.AddBundle(quoteGroup)
	*/
	r := mux.NewRouter()

	lst, err := kong.Whitelist(http.DefaultClient, secureUrl, "quote.submission.search", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
