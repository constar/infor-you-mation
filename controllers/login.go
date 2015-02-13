package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golang/glog"
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
	//sess := mc.StartSession()
	var user models.User
	inputs := mc.Input()
	user.Username = inputs.Get("username")
	user.Password = inputs.Get("password")
	err := models.ValidateUser(user)
	if err == nil {
		//sess.Set("username", user.Username)
		glog.Info("username:", user.Username, " login success!")
		mc.TplNames = "index.tpl"
	} else {
		glog.Error(err)
		mc.TplNames = "error.tpl"
	}
}
