package controllers;

import (

	b "gopkg.in/mgo.v2/bson"
	"github.com/astaxie/beego"
	"github.com/giter/hdmxx/models"
)


type NewSiteController struct {
	beego.Controller
}

func (this *NewSiteController) Get() {

	id := this.GetString("Id")

	s := &models.Site{}

	if id != "" {
		s,_ = models.GetSite(id)
	}

	this.Data["Site"] = s;
	this.Data["Users"],_ = models.ListUser()
	this.Layout = "layout.tpl"
	this.TplNames = "site.tpl"
}

type _s struct {

	Id string `form:"Id"`
	Name string `form:"Name"`
	Url string `form:"Url"`
	CheckPoint string `form:"CheckPoint"`
	Method string `form:"Method"`
	Duration int `form:"Duration"`
	Disabled bool `form:"Disabled"`
}

func (this *NewSiteController) Post(){

	s := _s{}

	if err := this.ParseForm(&s) ; err!= nil{
		this.Abort("401")
		return
	}

	var ss *models.Site
	
	if s.Id != "" {
		ss,_ = models.GetSite(s.Id)
	}else{
		ss = &models.Site{}
		ss.Id = b.NewObjectId()
	}

	ss.Name = (s.Name)
	ss.Url = (s.Url)
	ss.CheckPoint = (s.CheckPoint)
	ss.Method = (s.Method)
	ss.Duration = (s.Duration)
	ss.Disabled = (s.Disabled)

	if s.Id != "" {
		models.UpdateSite(*ss)
	}else{
		models.NewSite(*ss)
	}

	this.Redirect("/", 302)

}
