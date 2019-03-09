package routers

import (
	"github.com/astaxie/beego"
	"github.com/ziyoubiancheng/xcms/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
