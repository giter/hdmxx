package main

import (

	"strings"
	"time"
	_ "github.com/giter/hdmxx/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/giter/hdmxx/models"
)

func main() {

	go (func(){

		timer := time.NewTicker(5 * time.Second)

		for {
			select {
			case <-timer.C:
				models.DoSiteCheck()
			}
		}
	})();

	beego.SessionOn = true

	var AccessFilter = func(ctx *context.Context) {

		uri := ctx.Input.Uri()
		beego.Warn(uri)

		if strings.HasPrefix(uri, "/login.go") || strings.HasPrefix(uri, "/static/") || strings.HasPrefix(uri, "/favicon.ico") {
			return
		}

		u, ok := ctx.Input.Session("user").(*models.User)

		if !ok || u.Account == "" {
			ctx.Redirect(302, "/login.go")
			return
		}

		ctx.Input.SetData("GUser", ctx.Input.Session("user").(*models.User))
	}

	beego.InsertFilter("*",beego.BeforeRouter,AccessFilter)
	beego.Run()
}

