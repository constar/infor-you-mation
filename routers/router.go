package routers

import (
	"github.com/astaxie/beego"
	"github.com/yanyiwu/infor-you-mation/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/card", &controllers.CardController{})
}
