package main

import (
	_ "wxapi.credit/routers"

	"github.com/astaxie/beego"
	"wxapi.credit/services"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	services.ConfigInit()
	beego.Run()
}
