package controllers

import (
	. "ac/models"
)

type AddPetController struct {
	BaseController
}

func (u *AddPetController) Get() {
	u.TplName = "shangchuan.html"
}

func (u *AddPetController) Post() {
	Name := u.GetString("Name")
	Speci := u.GetString("Speci")
	Variety := u.GetString("Variety")
	Sex := u.GetIntWithDefault("Sex", -1)
	Age := u.GetIntWithDefault("Age", -1)
	Intro := u.GetString("Intro")
	Partner := u.GetIntWithDefault("Partner", -1)
	//PetImg := u.GetString("PetImg")

	if Sex == -1 || Age == -1 || Partner == -1 {
		u.Ctx.WriteString("Sex || Age || Partner is illegal")
		return
	}

	if Name == "" || Speci == "" || Variety == "" || Intro == "" {
		u.Ctx.WriteString("Name or Speci or Variety or Intro is blank")
		return
	}

	var pet Pet
	pet.Name = Name
	pet.Speci = Speci
	pet.Variety = Variety
	pet.Sex = Sex
	pet.Age = Age
	pet.Intro = Intro
	pet.Partner = Partner

	id, err := AddPet(pet)

	if err != nil {
		u.Data["json"] = map[string]interface{}{"code": 1, "message": "博客添加成功", "id": id}
	} else {
		u.Data["json"] = map[string]interface{}{"code": 0, "message": "博客添加出错"}
	}
	u.ServeJSON()
}
