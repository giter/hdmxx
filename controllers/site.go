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
	Type string `form:"Type"`
	CheckPoint string `form:"CheckPoint"`
	Method string `form:"Method"`
	Address string `form:"Address"`
	Port int `form:"Port"`
	Duration int `form:"Duration"`
	Disabled bool `form:"Disabled"`
	CTimeout int `form:"CTimeout"`
	RTimeout int `form:"RTimeout"`
	Input string `form:"Input"`
	Result string `form:"Result"`
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

	if u, ok := this.GetSession("user").(*models.User); ok && u != nil {

		found := false
		for _,user := range ss.Users {
			if user.Account == u.Account {
				found = true
				break
			}
		}

		if !found {
			ss.Users = append(ss.Users, *u)
		}
	}

	ss.Name = (s.Name)
	ss.Url = (s.Url)
	ss.Type = (s.Type)
	ss.CheckPoint = (s.CheckPoint)
	ss.Method = (s.Method)
	ss.Address = (s.Address)
	ss.Port = (s.Port)
	ss.Duration = (s.Duration)
	ss.Disabled = (s.Disabled)
	ss.CTimeout = (s.CTimeout)
	ss.RTimeout = (s.RTimeout)
	ss.Input = (s.Input)
	ss.Result = (s.Result)

	if s.Id != "" {
		models.UpdateSite(*ss)
	}else{
		models.NewSite(*ss)
	}

	this.Redirect(ROOT, 302)

}
