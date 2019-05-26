package controllers

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/ziyoubiancheng/xcms/consts"
	"github.com/ziyoubiancheng/xcms/models"
)

type MenuController struct {
	BaseController
}

func (c *MenuController) Index() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs.html"
	c.setTpl()
}

func (c *MenuController) List() {
	data, total := models.MenuList()
	c.listJsonResult(consts.JRCodeSucc, "ok", total, data)
}

func (c *MenuController) Add() {
	//选择父菜单数据
	data, _ := models.MenuList()
	var parentMenus []models.MenuModel
	for _, value := range data {
		if 0 == value.Parent {
			parentMenus = append(parentMenus, *value)
		}
	}
	c.Data["PMenus"] = parentMenus
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs_add.html"
	c.setTpl("menu/add.html", "common/layout_edit.html")
}
func (c *MenuController) AddDo() {
	var m models.MenuModel
	if err := c.ParseForm(&m); err == nil {
		id, _ := orm.NewOrm().Insert(&m)
		c.jsonResult(consts.JRCodeSucc, "ok", id)
	} else {
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
}

func (c *MenuController) Edit() {
	c.Data["Mid"] = c.GetString("mid")
	c.Data["Parent"], _ = c.GetInt("parent")
	c.Data["Name"] = c.GetString("name")
	c.Data["Seq"] = c.GetString("seq")

	//选择父菜单数据
	data, _ := models.MenuList()
	var parentMenus []models.MenuModel
	for _, value := range data {
		if 0 == value.Parent {
			parentMenus = append(parentMenus, *value)
		}
	}
	c.Data["PMenus"] = parentMenus

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs_edit.html"
	c.setTpl("menu/edit.html", "common/layout_edit.html")
}

func (c *MenuController) EditDo() {
	var m models.MenuModel
	if err := c.ParseForm(&m); err == nil {
		id, _ := orm.NewOrm().Update(&m)
		c.jsonResult(consts.JRCodeSucc, "ok", id)
	} else {
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
}

func (c *MenuController) DeleteDo() {
	if mid, err := c.GetInt("mid"); err == nil {
		num, _ := orm.NewOrm().Delete(&models.MenuModel{Mid: mid})
		c.jsonResult(consts.JRCodeSucc, "ok", num)
	} else {
		fmt.Println(err, mid)
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
}
