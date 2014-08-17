package routers

import (
	c "github.com/giter/hdmxx/controllers"
	"github.com/astaxie/beego"
)

func init() {

  beego.Router(c.ROOT, &c.MainController{})
  beego.Router(c.SITE, &c.NewSiteController{})
	beego.Router(c.LOGIN,&c.UserLogin{})
	beego.Router(c.LOGOUT,&c.UserLogin{})
}
