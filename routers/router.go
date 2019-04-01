package routers

import (
	"github.com/astaxie/beego"
	"github.com/ziyoubiancheng/xcms/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//menu
	beego.Router("/menu", &controllers.MenuController{}, "Get:Index")
	beego.Router("/menu/add", &controllers.MenuController{}, "Get:Add")
	beego.Router("/menu/adddo", &controllers.MenuController{}, "*:AddDo")
	beego.Router("/menu/edit", &controllers.MenuController{}, "Get:Edit")
	beego.Router("/menu/editdo", &controllers.MenuController{}, "*:EditDo")
	beego.Router("/menu/deletedo", &controllers.MenuController{}, "Get:DeleteDo")
	beego.Router("/menu/list", &controllers.MenuController{}, "*:List")
}
