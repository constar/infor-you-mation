package routers

import (
	"github.com/astaxie/beego"
	"inforyoumation/controllers"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
}
