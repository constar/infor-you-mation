package controllers

import (
	"../models"
	"github.com/astaxie/beego"
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
