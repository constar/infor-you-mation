package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yanyiwu/infor-you-mation/models"
	"strings"
)

type CardController struct {
	beego.Controller
}

func (this *CardController) Get() {
	k := this.GetString("k")
	//beego.Debug("CardController Get", k)
	this.Data["Card"] = models.GetCardByTopic(k, 20)
	this.Data["YesterdayAdd"] = models.GetYesterdayAddByKeyword(strings.ToLower(k))
	//beego.Debug(this.Data["Card"].(*models.Card))
	this.TplNames = "carddetail.html"
}

func (this *CardController) Post() {
	beego.Debug("CardController Post")
}
