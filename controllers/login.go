package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/lucky7ky/infor-you-mation/models"
	"github.com/yanyiwu/igo"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "login.html"
}

func (this *LoginController) Post() {
	inputs := this.Input()
	username := inputs.Get("username")
	passwd := inputs.Get("password")
	passwd = igo.GetMd5String(passwd)
	if models.ValidateUser(username, passwd) {
		glog.Info("username:", username, " login success!")
		this.TplNames = "index.html"
	} else {
		this.TplNames = "loginfailure.html"
	}
}
