package main

import (
	"flag"
	"github.com/astaxie/beego"
	_ "github.com/lucky7ky/infor-you-mation/routers"
)

func main() {
	flag.Parse()
	//beego.SessionOn = true
	beego.SetStaticPath("css", "views/css")
	beego.Run()
}
