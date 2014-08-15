package routers

import (
	"github.com/giter/hdmxx/controllers"
	"github.com/astaxie/beego"
)

func init() {

  beego.Router("/", &controllers.MainController{})
  beego.Router("/site.go", &controllers.NewSiteController{})
	beego.Router("/login.go",&controllers.UserLogin{})
	beego.Router("/logout.go",&controllers.UserLogin{})
}
