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
		beego.NSNamespace("/user",
			beego.NSRouter("/:id", &controller.User{}, "get:User"),
			beego.NSRouter("/:id/games", &controller.User{}, "get:Games"),
			beego.NSRouter("/:id/mostPlayedCharacter", &controller.User{}, "get:MostPlayedCharacter"),
			beego.NSRouter("/:id/character/:cn", &controller.User{}, "get:Character")))
}
