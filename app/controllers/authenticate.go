package controllers

import (
	"app/security"
	"github.com/beego/beego/v2/server/web"
)

type AuthenticateController struct {
	web.Controller
}

type LoginResponse struct {
	TokenString string `json:"jwt"`
}

func (c *AuthenticateController) Post() {
	tokenString := security.GenerateToken()
	security.Valid(tokenString)
	res := LoginResponse {tokenString}

	c.Data["json"] = &res
    c.ServeJSON()
}
