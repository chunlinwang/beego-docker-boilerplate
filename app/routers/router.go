package routers

import (
	"app/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
    web.Router("/", &controllers.MainController{})
}
