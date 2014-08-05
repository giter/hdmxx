package main

import (

	"time"
	_ "github.com/giter/hdmxx/routers"
	"github.com/astaxie/beego"
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

	beego.Run()
}

