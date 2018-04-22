package controllers

import (
	"actest/models"

	"github.com/astaxie/beego"
)

type Uploadcontrollers struct {
	BaseController
}

func (u *Uploadcontrollers) Get() {
	u.TplName = "shangchuan.html"
}

func (u *Uploadcontrollers) Post() {
	Name := u.GetString("Name")
	Speci := u.GetString("Speci")
	Variety := u.GetString("Variety")
	Sex := u.GetIntWithDefault("Sex", -1)
	Age := u.GetIntWithDefault("Age", -1)
	Intro := u.GetString("Intro")
	Partner := u.GetIntWithDefault("Partner", -1)
	PetImg := u.GetString("PetImg")

	if Sex == -1 || Age == -1 || Partner == -1 {
		u.Ctx.WriteString("Sex || Age || Partner is illegal")
		return
	}

	if Name == "" || Speci == "" || Variety == "" || Intro == "" {
		u.Ctx.WriteString("Name or Speci or Variety or Intro is blank")
		return
	}
	err := models.Pet(Name, Speci, Variety, Sex, Age, Intro, Partner, PetImg)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
}
