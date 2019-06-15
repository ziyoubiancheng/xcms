package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/ziyoubiancheng/xcms/consts"
	"github.com/ziyoubiancheng/xcms/models"
	"github.com/ziyoubiancheng/xcms/utils"
)

type UserController struct {
	BaseController
}

func (c *UserController) Index() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs.html"
	c.setTpl()
}

func (c *UserController) List() {
	page, err := c.GetInt("page")
	if err != nil {
		page = 1
	}
	size, err := c.GetInt("limit")
	if err != nil {
		size = 20
	}

	result, count := models.UserList(size, page)
	c.listJsonResult(consts.JRCodeSucc, "ok", count, result)
}

func (c *UserController) Add() {
	menu := models.ParentMenuList()
	menus := make(map[int]string)
	for _, v := range menu {
		menus[v.Mid] = v.Name
	}
	c.Data["Menus"] = menus
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs_add.html"
	c.setTpl("user/add.html", "common/layout_edit.html")
}
func (c *UserController) AddDo() {
	var m models.UserModel
	if err := c.ParseForm(&m); err == nil {
		//菜单权限
		menu := models.ParentMenuList()
		var menus []int
		for _, v := range menu {
			if "on" == c.GetString("userauth_"+strconv.Itoa(v.Mid)) {
				menus = append(menus, v.Mid)
			}
		}
		menujson, _ := json.Marshal(menus)
		m.AuthStr = string(menujson)
		m.Password = utils.Md5([]byte(m.Password))
		id, _ := orm.NewOrm().Insert(&m)
		c.jsonResult(consts.JRCodeSucc, "ok", id)
	} else {
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
}

func (c *UserController) Edit() {
	userid, _ := c.GetInt("userid")

	//初始化用户信息
	o := orm.NewOrm()
	var user = models.UserModel{UserId: userid}
	o.Read(&user)
	user.Password = ""
	c.Data["User"] = user

	authmap := make(map[int]bool)
	if len(user.AuthStr) > 0 {
		var authobj []int
		json.Unmarshal([]byte(user.AuthStr), &authobj)
		for _, v := range authobj {
			authmap[v] = true
		}
	}

	//初始化menu列表
	type Menuitem struct {
		Name    string
		Ischeck bool
	}
	menu := models.ParentMenuList()
	menus := make(map[int]Menuitem)
	for _, v := range menu {
		menus[v.Mid] = Menuitem{v.Name, authmap[v.Mid]}
	}
	c.Data["Menus"] = menus

	//页面设置
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs_edit.html"
	c.setTpl("user/edit.html", "common/layout_edit.html")
}

func (c *UserController) EditDo() {
	var m models.UserModel
	if err := c.ParseForm(&m); err == nil {
		//菜单权限
		menu := models.ParentMenuList()
		var menus []int
		for _, v := range menu {
			if "on" == c.GetString("userauth_"+strconv.Itoa(v.Mid)) {
				menus = append(menus, v.Mid)
			}
		}
		menujson, _ := json.Marshal(menus)
		m.AuthStr = string(menujson)
		//密码
		m.Password = utils.Md5([]byte(m.Password))
		id, _ := orm.NewOrm().Update(&m)
		c.jsonResult(consts.JRCodeSucc, "ok", id)
	} else {
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
}

func (c *UserController) DeleteDo() {
	if uid, err := c.GetInt("uid"); err == nil {
		num, _ := orm.NewOrm().Delete(&models.UserModel{UserId: uid})
		c.jsonResult(consts.JRCodeSucc, "1", num)
	} else {
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
}
