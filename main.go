package main

import (
	_ "Mybook/routers"
	"github.com/astaxie/beego"
	_ "Mybook/sysinit"
)

func main() {
	beego.Run()
}

