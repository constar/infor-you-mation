package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yanyiwu/infor-you-mation/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	cat := this.GetString("cat")
	switch cat {
	case "random":
		this.Data["CardFlows"] = models.GetRandomCardFlows(5)
	default:
		this.Data["CardFlows"] = models.GetCardFlows(5)
	}
	this.TplNames = "index.html"

}
