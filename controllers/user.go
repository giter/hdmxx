package controllers



import (

	"github.com/astaxie/beego"
	"github.com/giter/hdmxx/models"
)

type UserLogin struct {
	beego.Controller
}

func (this *UserLogin) Get() {

	this.Layout = "layout.tpl"
	this.TplNames = "login.tpl"
}

type _FUserLogin struct {

	Account string
	Password string
}

func (this *UserLogin) Post() {

	s := _FUserLogin{}

	if err := this.ParseForm(&s) ; err!= nil{

		this.Abort("401")
		return
	}

	u := models.UserLogin(s.Account,s.Password)

	if u.Account != "" {

		this.SetSession("user", u)
		this.Redirect("/",302)
	}

	this.Redirect("/login.go",302)
}

type UserLogout struct {
	beego.Controller
}

func (this *UserLogout) Get() {

	this.DelSession("user")
	this.Redirect("/",302)
}
