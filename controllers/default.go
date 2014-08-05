package controllers

import (

	"github.com/astaxie/beego"
	"github.com/giter/hdmxx/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	
	this.Data["Sites"] = models.ListSite()

	this.TplNames = "index.tpl"
}
