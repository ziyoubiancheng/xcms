package main

import (
	_ "github.com/ziyoubiancheng/xcms/routers"
	_ "github.com/ziyoubiancheng/xcms/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
