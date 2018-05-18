package main

import (
	"github.com/astaxie/beego/orm"
	_ "actest/routers"
	_ "actest/initial"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego"
	_"actest/models"
)


func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

