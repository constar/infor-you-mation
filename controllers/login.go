package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/yanyiwu/igo"
	"inforyoumation/models"
)

type LoginController struct {
	beego.Controller
}

func (mc *LoginController) Get() {
	mc.Data["Website"] = "InfoYouMation"
	mc.Data["Email"] = "i@yanyiwu.com"
	mc.TplNames = "login.tpl"
}

func (mc *LoginController) Post() {
	inputs := mc.Input()
	username := inputs.Get("username")
	passwd := inputs.Get("password")
	passwd = igo.GetMd5String(passwd)
	if models.ValidateUser(username, passwd) {
		glog.Info("username:", username, " login success!")
		mc.TplNames = "index.tpl"
	} else {
		mc.TplNames = "loginfailure.tpl"
	}
}
