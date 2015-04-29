package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/lucky7ky/infor-you-mation/models"
)

type TrendController struct {
	beego.Controller
}

func (this *TrendController) Get() {
	k := this.GetString("k")
	beego.Info("Trend Keyword:", k)
	cnts := models.GetTopicTrend(k)
	b, err := json.Marshal(cnts)
	if err != nil {
		beego.Error(err)
		return
	}
	beego.Info(string(b))
	this.Ctx.WriteString(string(b))
}

func (this *TrendController) Post() {
	this.Ctx.WriteString("unfinished")
}
