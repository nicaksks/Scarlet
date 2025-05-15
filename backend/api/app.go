package api

import (
	"scarlet/api/controller"
	"scarlet/api/loggers"
	"scarlet/api/middlewares"

	beego "github.com/beego/beego/v2/server/web"
)

func Start() {

	loggers.Start()

	beego.Router("*", &controller.NotFound{}, "*:Error")
	beego.AddNamespace(routes())

	//Change server port  in conf/app.conf
	beego.RunWithMiddleWares("", middlewares.RateLimit, middlewares.CorsAndCache, middlewares.ValidateRegion)
}

func routes() *beego.Namespace {
	return beego.NewNamespace("/v1",
		beego.NSNamespace("/profile",
			beego.NSRouter("/:id", &controller.Profile{}, "get:Profile"),
			beego.NSRouter("/:id/histories", &controller.Profile{}, "get:Histories"),
			beego.NSRouter("/:id/mostPlayedCharacter", &controller.Profile{}, "get:MostPlayedCharacter"),
			beego.NSRouter("/:id/character/:cn", &controller.Profile{}, "get:Character")))
}
