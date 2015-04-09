package controllers

import (
	"../models"
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/yanyiwu/igo"
)

type RegisterController struct {
	beego.Controller
}

func (rc *RegisterController) Get() {
	rc.TplNames = "register.html"
}

func (rc *RegisterController) Post() {
	inputs := rc.Input()
	username := inputs.Get("username")
	password := inputs.Get("password")
	checkpasswd := inputs.Get("checkpassword")
	if password != checkpasswd {
		glog.Errorf("password check failed [%s, %s]", password, checkpasswd)
		rc.TplNames = "registerfailure.html"
		return
	}
	password = igo.GetMd5String(password)
	glog.Info("username:", username, "password:", password)
	err := models.RegisterUser(username, password)
	if err == nil {
		rc.TplNames = "registersuccess.html"
	} else {
		glog.Error(err)
		rc.TplNames = "registerfailure.html"
	}
}
