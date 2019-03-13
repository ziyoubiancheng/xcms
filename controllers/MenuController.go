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

	c.jsonResult(consts.JRCodeSucc, "ok", menu)
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
	var m models.MenuModel
	if err := c.ParseForm(&m); err == nil {
		id, _ := orm.NewOrm().Update(&m)
		c.jsonResult(consts.JRCodeSucc, "ok", id)
	} else {
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}
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
