package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.Data["IsLogin"] = CheckAccount(this.Ctx)
	this.Data["topics"] = models.FindAllTopic()
	this.TplName = "home.html"
}
