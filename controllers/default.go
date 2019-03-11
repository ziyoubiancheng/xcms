package controllers

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//	orm.RegisterDataBase("default", "mysql", "root:135246@tcp(127.0.0.1:3307)/xcms?charset=utf8", 30)

	//	//	o := orm.NewOrm()

	//	//	var r RawSeter
	//	//r := o.Raw("INSERT INTO  xcms_formats (name) VALUES ('?');", "test2").Exec()
	//	fmt.Println("hello")

	db, err := sql.Open("mysql", "root:135246@tcp(127.0.0.1:3307)/xcms?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	rows, _ := db.Query("select mid,name from xcms_menu limit 10")
	mid := 0
	name := " "
	for rows.Next() {
		rows.Scan(&mid, &name)
		fmt.Println(mid, name)
	}

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
