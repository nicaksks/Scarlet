package controller

import (
	"net/http"
	"scarlet/api/errors"
	"scarlet/api/utils"
	"scarlet/internal/http/controller"
	"scarlet/internal/http/enum"

	"github.com/beego/beego/v2/server/web"
)

type Profile struct {
	web.Controller
	region       enum.Region
	userNum      string
	characterNum string
}

func (c *Profile) Prepare() {
	region := c.Ctx.Input.Query("r")

	c.region = utils.GetRegion[region]
	c.userNum = c.Ctx.Input.Param(":id")
	c.characterNum = c.Ctx.Input.Param(":cn")
}

func (c *Profile) Profile() {
	response, err := controller.Profile(c.region, c.userNum)

	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		c.Ctx.JSONResp(errors.ResponseError(http.StatusNotFound, err.Error()))
		return
	}

	c.Ctx.JSONResp(response)
}

func (c *Profile) Histories() {
	response, err := controller.Histories(c.region, c.userNum)

	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		c.Ctx.JSONResp(errors.ResponseError(http.StatusNotFound, err.Error()))
		return
	}

	c.Ctx.JSONResp(response)
}

func (c *Profile) MostPlayedCharacter() {
	response, err := controller.MostPlayedCharacter(c.region, c.userNum)

	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		c.Ctx.JSONResp(errors.ResponseError(http.StatusNotFound, err.Error()))
		return
	}

	c.Ctx.JSONResp(response)
}

func (c *Profile) Character() {
	response, err := controller.CharacterInfo(c.region, c.userNum, c.characterNum)

	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		c.Ctx.JSONResp(errors.ResponseError(http.StatusNotFound, err.Error()))
		return
	}

	c.Ctx.JSONResp(response)
}
