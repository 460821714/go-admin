package routers

import (
	"github.com/freewebsys/go-admin/go-admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
