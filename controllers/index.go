package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yanyiwu/infor-you-mation/models"
)

type IndexController struct {
	beego.Controller
}

func (mc *IndexController) Get() {
	mc.Data["CardFlows"] = models.GetCardFlows()
	mc.TplNames = "index.html"
}
