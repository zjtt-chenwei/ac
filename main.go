package main

import (
	_ "actest/routers"
	_ "actest/initial"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

