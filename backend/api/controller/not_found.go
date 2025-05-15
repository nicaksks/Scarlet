package controller

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

const BASE_URL = "scarlet-bs.fun"

type NotFound struct {
	web.Controller
}

func (c *NotFound) Error() {
	c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
	c.Ctx.JSONResp(map[string]interface{}{
		"Cod": 404,
		"Msg": "Play Black Survival: Project Lumia",
	})
}
