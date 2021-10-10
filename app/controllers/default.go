package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

type MainResponse struct {
    Website string `json:"website"`
    Email string `json:"email"`
}

func (c *MainController) Get() {
	res := MainResponse { "beego.me", "astaxie@gmail.com" }
	c.Data["json"] = &res
    c.ServeJSON()
}
