package main

import (
	"github.com/astaxie/beego"
	_ "infor-you-mation/routers"
)

func main() {
	//beego.SessionOn = true
	beego.Run()
}
