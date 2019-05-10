package main

import (
	"log"
	"os"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	"github.com/louisevanderlith/quote/core"
	"github.com/louisevanderlith/quote/routers"

	"github.com/astaxie/beego"
)

func main() {
	mode := os.Getenv("RUNMODE")
	pubPath := os.Getenv("KEYPATH")

	core.CreateContext()
	defer core.Shutdown()

	// Register with router
	name := beego.BConfig.AppName
	srv := mango.NewService(mode, name, pubPath, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
