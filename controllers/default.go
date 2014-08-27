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

	var u *models.User
	var ok bool

	if u, ok = this.GetSession("user").(*models.User); ok && u != nil {
		this.Data["GUser"] = u
	}else{
		this.Redirect(LOGIN,302)
		return
	}

	Query := b.M{}

	if !u.Admin {
		Query["Users.Account"] = u.Account
	}

	Stats := map[string]int {"ALL":0}

	var s []models.Site

	if err := models.FindSite(&Query).Select(b.M{"Type":1}).All(&s) ; err == nil {

		for _,k := range s {

			if _,ok := Stats[k.Type] ; !ok {
				Stats[k.Type] = 0
			}

			Stats[k.Type]++
			Stats["ALL"]++
		}
	}

	this.Data["Stats"] = Stats


	Query["Disabled"] = false
	Type := this.GetString("t")

	if Type != "" {
		Query["Type"] = Type
	}


	models.FindSite(Query).Sort("_id").Iter().All(&s)
	this.Data["Sites"] = b.M{"Name":"Active","Sites":s,"Color":"primary","Add":true}

	var d []models.Site
	Query["Disabled"] = true
	models.FindSite(Query).Sort("_id").Iter().All(&d)
	this.Data["Disabled"] = b.M{"Name":"Inactive","Sites":d,"Color":"warning","Add":false}

	this.Layout = "layout.tpl"
	this.TplNames = "index.tpl"
}

