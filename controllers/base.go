package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	isLogin bool
}

func (b *BaseController) Prepare() {
	userLogin := b.GetSession("userLogin")
	if userLogin == nil {
		b.isLogin = false
	} else {
		b.isLogin = true
	}
	b.Data["isLogin"] = b.isLogin
}

func (b *BaseController) Go404() {
	b.TplName = "404.html"
	return
}
