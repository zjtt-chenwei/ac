package controllers

import (
	"actest/models"

	"github.com/astaxie/beego"
)

type RegisterController struct {
	BaseController
}

func (r *RegisterController) Get() {
	r.TplName = "register.html"
}

func (r *RegisterController) Post() {
	uname := r.Input().Get("uname")
	pwd := r.Input().Get("pwd")

	err := models.Register(uname, pwd)

	if err != nil {
		beego.Error(err)
	}

	r.Redirect("/", 302)
}
