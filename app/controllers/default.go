package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	res := "Welcome!"
	c.Data["json"] = &res
    c.ServeJSON()
}
