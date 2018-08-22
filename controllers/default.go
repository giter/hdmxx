package controllers

import (
	"github.com/astaxie/beego"
	"github.com/giter/hdmxx/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	var u *models.User
	var ok bool

	if u, ok = this.GetSession("user").(*models.User); ok && u != nil {
		this.Data["GUser"] = u
	} else {
		this.Redirect(LOGIN, 302)
		return
	}

	this.Layout = "layout.tpl"
	this.TplName = "index.tpl"
}
