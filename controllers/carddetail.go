package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lucky7ky/infor-you-mation/models"
)

type CardDetailController struct {
	beego.Controller
}

func (this *CardDetailController) Get() {
	k := this.GetString("k")
	this.Data["jsonp"] = models.GetCardByTopic(k, CARDDETAIL_ROW_LIMIT)
	this.ServeJsonp()
}

func (this *CardDetailController) Post() {
}
