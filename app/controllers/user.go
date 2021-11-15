package controllers

import (
	"app/models"
	"fmt"
)

type UserController struct {
	BaseController
}

func (c *UserController) Get() {

	res := "Welcome!"
	c.Data["json"] = &res
  c.ServeJSON()
}


func (c *UserController) Post() {
	fmt.Println(c.Data)

	u := models.User{}
	if err := c.ParseForm(&u); err != nil {
			//handle error
	}
	
	fmt.Println(u)

	res := "Welcome!"
	c.Data["json"] = &res
  c.ServeJSON()
}
