package routers

import (
	"github.com/astaxie/beego"
	c "github.com/giter/hdmxx/controllers"
)

func init() {

	beego.Router(c.ROOT, &c.MainController{})
	beego.Router(c.LOGIN, &c.UserLogin{})
	beego.Router(c.LOGOUT, &c.UserLogin{})
}
