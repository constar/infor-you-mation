package controllers

import (
	"../models"
	"github.com/astaxie/beego"
)

const (
	ROW_LIMIT = 5
)

type CardListController struct {
	beego.Controller
}

func (this *CardListController) Get() {
	v := models.GetHotCardFlows(ROW_LIMIT)
	this.Data["jsonp"] = v
	this.ServeJsonp()
}
