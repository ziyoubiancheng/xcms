package routers

import (
	"github.com/astaxie/beego"
	"github.com/ziyoubiancheng/xcms/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//menu
	beego.Router("/menu", &controllers.MenuController{}, "Get:Index")
	beego.Router("/menu/add", &controllers.MenuController{}, "Post:Add")
	beego.Router("/menu/edit", &controllers.MenuController{}, "Get:Edit")
	beego.Router("/menu/delete", &controllers.MenuController{}, "Get:Delete")
	beego.Router("/menu/list", &controllers.MenuController{}, "*:List")
}
