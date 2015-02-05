package main

import (
	"flag"
	"github.com/astaxie/beego"
	_ "inforyoumation/routers"
)

func init() {
	flag.Parse()
}

func main() {
	//beego.SessionOn = true
	beego.Run()
}
