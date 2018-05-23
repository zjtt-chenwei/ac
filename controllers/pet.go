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

	var vmaps map[string]interface{}
	vmaps = make(map[string]interface{})

	condMap := make(map[string]string)
	condMap["speci"] = speci
	condMap["sex"] = sex
	condMap["variety"] = variety
	condMap["name"] = name

	speciMap := ShowValues("PetSpeci")
	variMap := ShowValues("PetVari")

	vmaps["speci"] = speciMap
	vmaps["vari"] = variMap

	ap.Data["json"] = vmaps
	ap.ServeJSON()

	ap.TplName = "addpet.html"
}
