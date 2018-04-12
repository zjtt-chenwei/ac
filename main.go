package main

import (
	_ "actest/routers"
	_ "actest/initial"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

