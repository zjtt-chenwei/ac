package controllers

import (
	"actest/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

type UploadController struct {
	BaseController
}

// @Title UploadFile
// @Description 上传文件(单个)
// @Param imgFile formData file true "上传的文件"
// @Success 200 {string} 上传成功
// @Failture 404 上传失败

func (upl *UploadController) Post() {
	f, h, err := upl.GetFile("imgFile")

	defer f.Close()

	// 创建上传路径
	now := time.Now()
	dir := "./static/uploadfile/" + strconv.Itoa(now.Year()) + "/" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())
	err1 := os.MkdirAll(dir, 0755)
	if err1 != nil {
		upl.Data["json"] = map[string]interface{}{"error": 1, "message": "目录权限不够"}
		upl.ServeJSON()
		return
	}

	filename := h.Filename
	ext := utils.SubString(filename, strings.LastIndex(filename, "."), 5)
	filename = utils.GetGuid() + ext

	if err != nil {
		upl.Data["json"] = map[string]interface{}{"error": 1, "message": err}
	} else {
		upl.SaveToFile("imgFile", dir+"/"+filename)
		upl.Data["json"] = map[string]interface{}{"error": 0, "url": strings.Replace(dir, ".", "", 1) + "/" + filename}
	}
	upl.ServeJSON()
}
