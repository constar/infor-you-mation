package routers

import (
	"..//controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/card", &controllers.CardController{})
	beego.Router("/trend", &controllers.TrendController{})
}
