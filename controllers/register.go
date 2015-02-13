package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/yanyiwu/igo"
	"inforyoumation/models"
)

type RegisterController struct {
	beego.Controller
}

func (rc *RegisterController) Get() {
	rc.TplNames = "register.tpl"
}

func (rc *RegisterController) Post() {
	inputs := rc.Input()
	username := inputs.Get("username")
	password := inputs.Get("password")
	checkpasswd := inputs.Get("checkpassword")
	if password != checkpasswd {
		glog.Errorf("password check failed [%s, %s]", password, checkpasswd)
		rc.TplNames = "registerfailure.tpl"
		return
	}
	password = igo.GetMd5String(password)
	glog.Info("username:", username, "password:", password)
	err := models.RegisterUser(username, password)
	if err == nil {
		rc.TplNames = "registersuccess.tpl"
	} else {
		glog.Error(err)
		rc.TplNames = "registerfailure.tpl"
	}
}
