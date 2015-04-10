package controllers

import (
	"../models"
	"github.com/astaxie/beego"
)

type CardListController struct {
	beego.Controller
}

func (this *CardListController) Get() {
	this.Data["jsonp"] = models.GetHotCardFlows(ROW_LIMIT)
	this.ServeJsonp()
}
