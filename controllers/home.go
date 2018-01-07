package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (h *HomeController) Get() {
	h.Data["IsHome"] = true
	h.TplName = "home.html"
	h.Data["IsLogin"] = checkAccount(h.Ctx)
}
 