package main

import (
	"github.com/astaxie/beego"
	_ "inforyoumation/routers"
)

func main() {
	//beego.SessionOn = true
	beego.Run()
}
