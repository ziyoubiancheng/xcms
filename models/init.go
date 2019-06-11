package models

import (
	//"strconv"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(new(MenuModel), new(UserModel), new(DataModel))
}

func TbNameMenu() string {
	return "xcms_menu"
}

func TbNameUser() string {
	return "xcms_user"
}

func TbNameData() string {
	return "xcms_data"
}
