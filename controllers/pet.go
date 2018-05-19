package controllers

// import (
// 	. "actest/models"
// )

type AddPetController struct {
	BaseController
}

func (ap *AddPetController) Get() {
	ap.Data["IsAddpet"] = true
	// speci := ap.GetString("speci")
	// variety := ap.GetString("variety")
	// sex := ap.GetString("sex")
	// name := ap.GetString("name")

	// condMap := make(map[string]string)
	// condMap["speci"] = speci
	// condMap["sex"] = sex
	// condMap["variety"] = variety
	// condMap["name"] = name

	// _, _, specis := DistinctNum("pet_speci", "", "", "speci")
	// _, _, varis := DistinctNum("pet_vari", "pet_speci", speci, "speci")

	// ap.Data["specis"] = specis
	// ap.Data["varis"] = varis

	ap.TplName = "addpet.html"
}
