package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yanyiwu/infor-you-mation/models"
)

type IndexController struct {
	beego.Controller
}

const (
	ROW_LIMIT = 5
)

func (this *IndexController) Get() {
	cat := this.GetString("cat")
	switch cat {
	case "random":
		this.Data["CardFlows"] = models.GetRandomCardFlows(ROW_LIMIT)
	case "hot":
		this.Data["CardFlows"] = models.GetHotCardFlows(ROW_LIMIT)
	default:
		this.Data["CardFlows"] = models.GetCardFlows(ROW_LIMIT)
	}
	this.TplNames = "index.html"

}
