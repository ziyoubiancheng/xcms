package models

import (
	"strconv"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(new(MenuModel), new(UserModel))
}

func TbNameMenu() string {
	return "xcms_menu"
}

func TbNameUser() string {
	return "xcms_user"
}

func TbNameData(mid int) string {
	return "xcms_data_" + strconv.Itoa(mid)
}
