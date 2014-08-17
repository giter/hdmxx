package controllers

import (

	b "gopkg.in/mgo.v2/bson"
	"github.com/astaxie/beego"
	"github.com/giter/hdmxx/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	if u, ok := this.GetSession("user").(models.User); ok {
		this.Data["GUser"] = u
	}else{
		this.Redirect(LOGIN,302)
	}


	//SiteColl().Find(nil).Sort("_id").Iter().All(&result)
	var d []models.Site
	models.FindSite(b.M{"Disabled":true}).Sort("_id").Iter().All(&d)

	this.Data["Disabled"] = b.M{"Name":"Disabled","Sites":d,"Color":"warning"}

	var s []models.Site
	models.FindSite(b.M{"Disabled":false}).Sort("_id").Iter().All(&s)

	this.Data["Sites"] = b.M{"Name":"Active","Sites":s,"Color":"success"}


	this.Layout = "layout.tpl"
	this.TplNames = "index.tpl"
}

