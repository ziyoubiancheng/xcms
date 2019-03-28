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

	c.setTpl()
}

func (c *MenuController) List() {
	var m models.MenuModel

	c.listJsonResult(consts.JRCodeSucc, "ok", 20, m.List())
}

func (c *MenuController) Add() {
	var m models.MenuModel
	if err := c.ParseForm(&m); err == nil {
		id, _ := orm.NewOrm().Insert(&m)
		c.jsonResult(consts.JRCodeSucc, "ok", id)
	} else {
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
}

func (c *MenuController) Edit() {
	//	var m models.MenuModel
	//	if err := c.ParseForm(&m); err == nil {
	//		id, _ := orm.NewOrm().Update(&m)
	//		c.jsonResult(consts.JRCodeSucc, "ok", id)
	//	} else {
	//		c.jsonResult(consts.JRCodeFailed, "", 0)
	//	}
	c.jsonResult(consts.JRCodeSucc, "ok", "edit")
}

func (c *MenuController) Delete() {
	if mid, err := c.GetInt("mid"); err == nil {
		num, _ := orm.NewOrm().Delete(&models.MenuModel{Mid: mid})
		c.jsonResult(consts.JRCodeSucc, "1", num)
	} else {
		fmt.Println(err, mid)
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
}
