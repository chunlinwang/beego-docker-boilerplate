package controllers

import (
	// "app/security"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"

	// "strings"
)

const (
	Authorization  = "Authorization"
)

type BaseController struct {
	web.Controller
}

func (c *BaseController) Prepare() {
	logs.Info("user", c)

	// bearerToken := c.Ctx.Request.Header.Get(Authorization)
	// token := strings.Fields(bearerToken)
	// claims, isValid := security.Valid(token[1])

	// if !isValid {
	// 	c.Data["json"] = map[string]interface{}{"error": "jwt is not valid"}
	// 	c.ServeJSON()
	// 	c.StopRun()
	// }

	// c.Data["iss"] = claims["iss"]
}

func (c *BaseController) Finish() {
	c.Ctx.ResponseWriter.ResponseWriter.Header().Add("responseTime", string(c.Ctx.ResponseWriter.Elapsed))
}
