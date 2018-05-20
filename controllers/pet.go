package controllers

import (
	. "actest/models"
)

type AddPetController struct {
	BaseController
}

func (ap *AddPetController) Get() {
	ap.Data["IsAddpet"] = true
	speci := ap.GetString("speci")
	variety := ap.GetString("variety")
	sex := ap.GetString("sex")
	name := ap.GetString("name")

	condMap := make(map[string]string)
	condMap["speci"] = speci
	condMap["sex"] = sex
	condMap["variety"] = variety
	condMap["name"] = name
	speci = "猫"

	_, _, specis := ShowValues("PetSpeci", "", "", "Name")
	_, _, varis := ShowValues("PetVari", "PetSpeci", speci, "Name")

	ap.Data["specis"] = specis
	ap.Data["varis"] = varis

	ap.TplName = "addpet.html"
}
