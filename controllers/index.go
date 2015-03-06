package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yanyiwu/infor-you-mation/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["CardFlows"] = models.GetCardFlows(5)
	this.TplNames = "index.html"
}
