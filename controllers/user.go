package controllers

import (
	"actest/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

func (lin *LoginController) Get() {
	Login := lin.isLogin
	if Login {
		lin.Redirect("/", 302)
	} else {
		lin.TplName = "login.html"
	}
}

func (lin *LoginController) Post() {
	uname := lin.GetString("uname")
	pwd := lin.GetString("pwd")
	// autoLogin := lin.GetString("autoLogin") == "on"

	if "" == uname {
		lin.Data["json"] = map[string]interface{}{"code": 0, "message": "账号不能为空"}
		lin.ServeJSON()
	}
	if "" == pwd {
		lin.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		lin.ServeJSON()
	}

	err, user := models.CheckLog(uname, pwd)

	if err == nil {
		lin.SetSession("userLogin", "1")
		lin.SetSession("uname",uname)
		lin.Data["json"] = map[string]interface{}{"code": 1, "message": "贺喜你，登录成功", "user": user}
	} else {
		lin.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	lin.ServeJSON()
}

type LogoutController struct {
	BaseController
}

func (lout *LogoutController) Get() {
	lout.DelSession("userLogin")
	lout.DelSession("uname")
	lout.Redirect("/", 302)
}

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


// func (lin *LoginController) Get() {
// 	isExit := lin.Input().Get("exit") == "true"
// 	if isExit {
// 		lin.Ctx.SetCookie("uname", "", -1, "/")
// 		lin.Ctx.SetCookie("pwd", "", -1, "/")
// 		lin.Redirect("/", 302)
// 		return
// 	}
// 	lin.TplName = "login.html"
// }

// func (lin *LoginController) Post() {
// 	uname := lin.Input().Get("uname")
// 	pwd := lin.Input().Get("pwd")
// 	autoLogin := lin.Input().Get("autoLogin") == "on"
// 	err := models.CheckLog(uname, pwd)
// 	if err == nil {
// 		maxAge := 0
// 		if autoLogin {
// 			maxAge = 1<<31 - 1
// 		}
// 		lin.Ctx.SetCookie("uname", uname, maxAge, "/")
// 		lin.Ctx.SetCookie("pwd", pwd, maxAge, "/")

// 	}
// 	lin.Redirect("/", 302)
// 	return

// }

// func checkAccount(ctx *context.Context) bool {
// 	ck, err := ctx.Request.Cookie("uname")
// 	if err != nil {
// 		return false
// 	}

// 	uname := ck.Value

// 	ck, err = ctx.Request.Cookie("pwd")
// 	if err != nil {
// 		return false
// 	}
// 	pwd := ck.Value

// 	err = models.CheckLog(uname, pwd)
// 	if err == nil {
// 		return true
// 	} else {
// 		return false
// 	}
// }