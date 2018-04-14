package controllers

type PetmatchController struct {
	BaseController
}

func (pm *PetmatchController) Get() {
	pm.Data["IsPetmatch"] = true
	pm.TplName = "petmatch.html"
}
