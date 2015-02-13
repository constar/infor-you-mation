package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"inforyoumation/models"
)

type RegisterController struct {
	beego.Controller
}

func (rc *RegisterController) Get() {
	rc.TplNames = "register.tpl"
}

func (rc *RegisterController) Post() {
	var user models.User
	inputs := rc.Input()
	user.Username = inputs.Get("username")
	user.Password = inputs.Get("password")
	glog.Info("user:", user.Username, "password:", user.Password)
	err := models.RegisterUser(user.Username, user.Password)
	if err == nil {
		rc.TplNames = "registersuccess.tpl"
	} else {
		glog.Error(err)
		rc.TplNames = "registerfailure.tpl"
	}
}
