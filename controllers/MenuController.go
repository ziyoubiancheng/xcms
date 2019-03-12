package controllers

import (
	"encoding/json"
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

	for _, v := range rows { //查询出来的数组
		fmt.Println(v.Mid, v.Parent, v.Name)
	}

	c.jsonResult(consts.JRCodeSucc, "ok", rows)
}

func (c *MenuController) Add() {
	var m models.MenuModel
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &m); err == nil {
		id, _ := orm.NewOrm().Insert(&m)
		c.jsonResult(consts.JRCodeSucc, "ok", id)
	} else {
		c.jsonResult(consts.JRCodeFailed, "", 0)
	}

	//id, _ := orm.NewOrm().Insert(models.MenuModel{Name: "hello", Parent: 1})
	//c.jsonResult(consts.JRCodeSucc, "1", id)
}

func (c *MenuController) Delete() {
	num, _ := orm.NewOrm().Delete(9)

	c.jsonResult(consts.JRCodeSucc, "1", num)
}

func (c *MenuController) Edit() {
	num, _ := orm.NewOrm().Update(models.MenuModel{Mid: 5, Parent: 2}, "Parent")

	c.jsonResult(consts.JRCodeSucc, "1", num)
}
