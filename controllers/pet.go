package controllers

import (
	. "actest/models"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type AddPetController struct {
	BaseController
}

func (ap *AddPetController) Get() {
	ap.TplName = "addpet.html"
}

func (ap *AddPetController) Post() {
	ap.Data["IsAddpet"] = true
	speci := ap.GetString("speci")
	variety := ap.GetString("variety")
	sexstr := ap.GetString("sex")
	name := ap.GetString("name")
	intro := ap.GetString("intro")
	birthString := ap.GetString("birth")
	partnerstr := ap.GetString("partner")
	
	sex, _ := strconv.ParseBool(sexstr)
	partner, _ := strconv.ParseBool(partnerstr)

	f, h, err := ap.GetFile("imgFile")
	defer f.Close()
	if err != nil {
		ap.Data["jsonErr"] = map[string]interface{}{"error": 1, "message": err}
	}
	dir, filename := UploadFile(h)
	fileURL := dir + "/" + filename
	ap.SaveToFile("imgFile", fileURL)
	ap.Data["jsonErr2"] = map[string]interface{}{"error": 0, "url": strings.Replace(dir, ".", "", 1) + "/" + filename}

	birth, _ := time.Parse(birthString, "2000-01-01")

	var newpetimg PetImg
	var newpet Pet
	newpetimg.ImgURL = fileURL
	newpetimg.Cover = true

	newpet.Sex = sex
	newpet.Variety = variety
	newpet.Speci = speci
	newpet.Partner = partner
	newpet.Intro = intro
	newpet.Birth = birth
	newpet.Name = name

	err1 := AddPetInfo(newpet, newpetimg)
	if err1 != nil {

	}

	var vmaps map[string]interface{}
	vmaps = make(map[string]interface{})

	// condMap := make(map[string]string)

	speciMap := ShowValues("PetSpeci")
	variMap := ShowValues("PetVari")

	vmaps["speci"] = speciMap
	vmaps["vari"] = variMap

	jsons, errs := json.Marshal(vmaps)
	if errs != nil {
	}

	ap.Data["json"] = string(jsons)
	ap.Data["specis"] = speciMap
	// ap.ServeJSON()

	ap.TplName = "addpet.html"
}
