package controllers

import (
	"../models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

const (
	ROW_LIMIT = 5
)

func (this *IndexController) Get() {
	cb := this.GetString("callback")
	beego.Info("callback", cb)
	if cb == "gethotcardflow" {
		b, err := json.Marshal(models.GetHotCardFlows(ROW_LIMIT))
		if err != nil {
			beego.Error(err)
			return
		}
		s := cb + "(" + string(b) + ");"
		this.Ctx.WriteString(s)
	} else {
		this.Data["CardFlows"] = models.GetHotCardFlows(ROW_LIMIT)
		this.TplNames = "index.html"
	}
}
