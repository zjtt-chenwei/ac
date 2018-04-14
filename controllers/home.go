package controllers

type HomeController struct {
	BaseController
}

func (h *HomeController) Get() {
	h.Data["IsHome"] = true
	h.TplName = "home.html"
	// h.Data["IsLogin"] = checkAccount(h.Ctx)
}
