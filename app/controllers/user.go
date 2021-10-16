package controllers

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	res := "Welcome!"
	c.Data["json"] = &res
    c.ServeJSON()
}
