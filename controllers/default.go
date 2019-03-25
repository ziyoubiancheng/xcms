package controllers

import (
	//	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
	//	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/ziyoubiancheng/xcms/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//	//	orm.RegisterDataBase("default", "mysql", "root:135246@tcp(127.0.0.1:3307)/xcms?charset=utf8", 30)

	//	//	//	o := orm.NewOrm()

	//	//	//	var r RawSeter
	//	//	//r := o.Raw("INSERT INTO  xcms_formats (name) VALUES ('?');", "test2").Exec()
	//	//	fmt.Println("hello")

	//	db, err := sql.Open("mysql", "root:135246@tcp(127.0.0.1:3307)/xcms?charset=utf8")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	defer db.Close()
	//	rows, _ := db.Query("select mid,name from xcms_menu limit 10")
	//	mid := 0
	//	name := " "
	//	for rows.Next() {
	//		rows.Scan(&mid, &name)
	//		fmt.Println(mid, name)
	//	}

	//	c.Data["Website"] = "beego.me"
	//	c.Data["Email"] = "astaxie@gmail.com"
	//	c.TplName = "index.tpl"
	c.Data["Test"] = "test1"

	var m models.MenuModel
	rows := m.List()

	var menu = make(map[int]models.MenuTree)
	for _, v := range rows { //查询出来的数组
		//fmt.Println(v.Mid, v.Parent, v.Name)
		if 0 == v.Mtype {
			var tree = new(models.MenuTree)
			tree.MenuModel = *v
			menu[v.Mid] = *tree
		} else {
			if tmp, ok := menu[v.Parent]; ok {
				tmp.Child = append(tmp.Child, *v)
				menu[v.Parent] = tmp
			}
		}
	}
	fmt.Println(menu)
	//c.Data["menu"] = "{\"code\": 0,\"msg\": \"ok\",\"data\": {  \"1\": {    \"Mid\": 1,    \"Mtype\": 0,    \"Parent\": 0,    \"Seq\": 0,    \"Name\": \"系统菜单\",    \"Fid\": 0,    \"Role\": 0,    \"Child\": [      {        \"Mid\": 4,        \"Mtype\": 10,        \"Parent\": 1,        \"Seq\": 5,        \"Name\": \"角色管理\",        \"Fid\": 0,        \"Role\": 0      },      {        \"Mid\": 3,        \"Mtype\": 10,        \"Parent\": 1,        \"Seq\": 4,        \"Name\": \"用户管理\",        \"Fid\": 0,        \"Role\": 0      },      {        \"Mid\": 5,        \"Mtype\": 10,        \"Parent\": 1,        \"Seq\": 3,        \"Name\": \"菜单管理\",        \"Fid\": 0,        \"Role\": 0      }    ]  }}}"
	c.Data["Menu"] = menu
	c.Layout = "common/layout.html"
	c.TplName = "testlay.html"
}
