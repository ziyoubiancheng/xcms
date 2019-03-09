package controllers

import (
	"fmt"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	orm.RegisterDataBase("default", "mysql", "root:135246@tcp(127.0.0.1:3307)/xcms?charset=utf8", 30)

	//	o := orm.NewOrm()

	//	var r RawSeter
	//r := o.Raw("INSERT INTO  xcms_formats (name) VALUES ('?');", "test2").Exec()
	fmt.Println("hello")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
