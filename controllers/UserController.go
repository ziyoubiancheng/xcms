package controllers

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/ziyoubiancheng/xcms/consts"
	"github.com/ziyoubiancheng/xcms/models"
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
	size, err := c.GetInt("size")
	if err != nil {
		size = 20
	}

	result, count := models.UserList(size, page)
	c.listJsonResult(consts.JRCodeSucc, "ok", count, result)
}

func (c *UserController) Add() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs_add.html"
	c.setTpl("user/add.html", "common/layout_edit.html")
}
func (c UserController) AddDo() {
	var m models.UserModel
	if err := c.ParseForm(&m); err == nil {
		id, _ := orm.NewOrm().Insert(&m)
		fmt.Println(m)
		c.jsonResult(consts.JRCodeSucc, "ok", id)
	} else {
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
}

func (c UserController) Edit() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "user/footerjs_edit.html"
	c.setTpl("user/edit.html", "common/layout_edit.html")
}
func (c UserController) EditDo() {
	var m models.UserModel
	if err := c.ParseForm(&m); err == nil {
		id, _ := orm.NewOrm().Update(&m)
		fmt.Println(m)
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
		fmt.Println(err, uid)
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
}
