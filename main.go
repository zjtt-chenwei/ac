package main

import (
	"github.com/astaxie/beego/orm"
	_ "actest/routers"
	_ "actest/initial"
	"github.com/astaxie/beego"
	_"actest/models"
)


func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

