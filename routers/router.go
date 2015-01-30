package routers

import (
	"github.com/astaxie/beego"
	"infor-you-mation/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/register", &controllers.RegisterController{})
}
