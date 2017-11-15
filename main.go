package main

import (
	_ "github.com/arifikhsan/bee/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/auth"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, auth.Basic("username", "secretpassword"))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
	}
	beego.Run()
}
