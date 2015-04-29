package routers

import (
	"..//controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/cardlist", &controllers.CardListController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/carddetail", &controllers.CardDetailController{})
	beego.Router("/trend", &controllers.TrendController{})
}
