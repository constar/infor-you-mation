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
	user.Pwd = inputs.Get("pwd")
	err := models.RegisterUser(user)
	if err == nil {
		rc.TplNames = "success.tpl"
	} else {
		glog.Error(err)
		rc.TplNames = "error.tpl"
	}
}
