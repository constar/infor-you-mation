package routers

import (
	"github.com/astaxie/beego"
	"github.com/yanyiwu/infor-you-mation/controllers"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
}
