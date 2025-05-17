package controller

import (
	"net/http"
	"scarlet/api/errors"
	"scarlet/api/utils"
	"scarlet/internal/http/controller"
	"scarlet/internal/http/enum"

	"github.com/beego/beego/v2/server/web"
)

type User struct {
	web.Controller
	region       enum.Region
	userNum      string
	characterNum string
}

func (c *User) Prepare() {
	region := c.Ctx.Input.Query("r")

	c.region = utils.GetRegion[region]
	c.userNum = c.Ctx.Input.Param(":id")
	c.characterNum = c.Ctx.Input.Param(":cn")
}

func (c *User) User() {
	response, err := controller.User(c.region, c.userNum)

	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		c.Ctx.JSONResp(errors.ResponseError(http.StatusNotFound, err.Error()))
		return
	}

	c.Ctx.JSONResp(response)
}

func (c *User) Games() {
	response, err := controller.Games(c.region, c.userNum)

	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		c.Ctx.JSONResp(errors.ResponseError(http.StatusNotFound, err.Error()))
		return
	}

	c.Ctx.JSONResp(response)
}

func (c *User) MostPlayedCharacter() {
	response, err := controller.MostPlayedCharacter(c.region, c.userNum)

	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		c.Ctx.JSONResp(errors.ResponseError(http.StatusNotFound, err.Error()))
		return
	}

	c.Ctx.JSONResp(response)
}

func (c *User) Character() {
	response, err := controller.CharacterInfo(c.region, c.userNum, c.characterNum)

	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		c.Ctx.JSONResp(errors.ResponseError(http.StatusNotFound, err.Error()))
		return
	}

	c.Ctx.JSONResp(response)
}
