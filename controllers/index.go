package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"inforyoumation/models"
)

type MainController struct {
	beego.Controller
}

func (mc *MainController) Get() {
	mc.Data["Website"] = "InfoYouMation"
	mc.Data["Email"] = "i@yanyiwu.com"
	mc.TplNames = "index.tpl"
}

func (mc *MainController) Post() {
	sess := mc.StartSession()
	var user models.User
	inputs := mc.Input()
	user.Username = inputs.Get("username")
	user.Pwd = inputs.Get("pwd")
	err := models.ValidateUser(user)
	if err == nil {
		sess.Set("username", user.Username)
		fmt.Println("username:", sess.Get("username"))
		mc.TplNames = "seccess.tpl"
	} else {
		fmt.Println(err)
		mc.TplNames = "error.tpl"
	}
}
