package main

import (
	"flag"
	"github.com/astaxie/beego"
	_ "inforyoumation/routers"
)

func main() {
	flag.Parse()
	//beego.SessionOn = true
	//beego.SetStaticPath("/publish", "statics")
	beego.Run()
}
