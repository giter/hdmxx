package routers

import (
	"github.com/giter/hdmxx/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
