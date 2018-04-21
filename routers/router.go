package routers

import (
	"actest/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/petmatch", &controllers.ListPetController{})
	beego.Router("/404.html", &controllers.BaseController{}, "*:Go404")
}
