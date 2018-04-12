package controllers

import (
	"github.com/astaxie/beego"
)

type PetmatchController struct {
	beego.Controller
}

func (pm *PetmatchController) Get() {
	pm.Data["IsPetmatch"] = true
	pm.TplName = "petmatch.html"
	// pm.Data["IsLogin"] = checkAccount(pm.Ctx)
}

