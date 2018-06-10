package controllers

import (
	."actest/models"

	"github.com/astaxie/beego"
)


type RegisterController struct {
	BaseController
}

type AccountController struct {
	BaseController
}

func (ac *AccountController) Get() {
	Login := ac.isLogin
	if Login {
		ac.Redirect("/", 302)
	} else {
		ac.TplName = "login.html"
	}
}

func (ac *AccountController) Post() {
	action := ac.GetString("action")
	if action == "login"{
		uname := ac.GetString("uname")
		pwd := ac.GetString("pwd")
		// autoLogin := ac.GetString("autoLogin") == "on"

		if "" == uname {
			ac.Data["json"] = map[string]interface{}{"code": 0, "message": "账号不能为空"}
			ac.ServeJSON()
		}
		if "" == pwd {
			ac.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
			ac.ServeJSON()
		}

		err, user := CheckLog(uname, pwd)

		if err == nil {
			ac.SetSession("userLogin", "1")
			ac.SetSession("uname",uname)
			ac.Data["json"] = map[string]interface{}{"code": 1, "message": "贺喜你，登录成功", "user": user}
		} else {
			ac.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
		}
		ac.ServeJSON()
	}else if action == "regist"{
		// message := ac.GetString("message")
		uname := ac.GetString("uname")
		pwd := ac.GetString("pwd")

		err := Register(uname, pwd)

		if err != nil {
			beego.Error(err)
		}

		ac.Redirect("/", 302)
	}
}

type LogoutController struct {
	BaseController
}

func (lout *LogoutController) Get() {
	lout.DelSession("userLogin")
	lout.DelSession("uname")
	lout.Redirect("/", 302)
}



// func (ac *AccountController) Get() {
// 	isExit := ac.Input().Get("exit") == "true"
// 	if isExit {
// 		ac.Ctx.SetCookie("uname", "", -1, "/")
// 		ac.Ctx.SetCookie("pwd", "", -1, "/")
// 		ac.Redirect("/", 302)
// 		return
// 	}
// 	ac.TplName = "login.html"
// }

// func (ac *AccountController) Post() {
// 	uname := ac.Input().Get("uname")
// 	pwd := ac.Input().Get("pwd")
// 	autoLogin := ac.Input().Get("autoLogin") == "on"
// 	err := models.CheckLog(uname, pwd)
// 	if err == nil {
// 		maxAge := 0
// 		if autoLogin {
// 			maxAge = 1<<31 - 1
// 		}
// 		ac.Ctx.SetCookie("uname", uname, maxAge, "/")
// 		ac.Ctx.SetCookie("pwd", pwd, maxAge, "/")

// 	}
// 	ac.Redirect("/", 302)
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