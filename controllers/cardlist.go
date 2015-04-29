package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lucky7ky/infor-you-mation/models"
)

type CardListController struct {
	beego.Controller
}

func (this *CardListController) Get() {
	this.Data["jsonp"] = models.GetHotCardFlows(ROW_LIMIT)
	this.ServeJsonp()
}
