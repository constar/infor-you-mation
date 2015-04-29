package main

import (
	_ "./routers"
	"flag"
	"github.com/astaxie/beego"
)

func main() {
	flag.Parse()
	//beego.SessionOn = true
	beego.SetStaticPath("css", "views/css")
	beego.Run()
}
