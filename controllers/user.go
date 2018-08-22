package controllers

import (
	"crypto/md5"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/giter/hdmxx/models"
)

type UserLogin struct {
	beego.Controller
}

func (this *UserLogin) Get() {

	this.Layout = "layout.tpl"
	this.TplName = "login.tpl"
}

type FUserLogin struct {
	Account  string
	Password string
}

func (this *UserLogin) Post() {

	s := FUserLogin{}

	if err := this.ParseForm(&s); err != nil {

		this.Abort("401")
		return
	}

	s.Password = fmt.Sprintf("%x", md5.Sum([]byte(s.Password)))

	u, err := models.NewUserService().Login(s.Account, s.Password)

	if err != nil {
		beego.Error(err)
		this.Redirect(LOGIN, 302)
	}

	if u.Account != "" {

		this.SetSession("user", u)
		this.Redirect(ROOT, 302)
	} else {
		this.Redirect(LOGIN, 302)
	}
}

type UserLogout struct {
	beego.Controller
}

func (this *UserLogout) Get() {

	this.DelSession("user")
	this.Redirect(ROOT, 302)
}
