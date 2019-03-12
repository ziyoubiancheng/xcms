package controllers

import (
	"fmt"

	"github.com/ziyoubiancheng/xcms/consts"
	"github.com/ziyoubiancheng/xcms/models"
)

type MenuController struct {
	BaseController
}

func (c *MenuController) Index() {
	var m models.MenuModel
	rows, _ := m.List()

	menus := make(map[string]string)
	len := len(rows)
	for k, v := range rows { //查询出来的数组
		fmt.Println(k, v)
		fmt.Println(v["mtype"])
		menus[v["mtype"]] = "hello"
	}
	fmt.Println(len)

	c.jsonResult(consts.JRCodeSucc, "ok", menus)
}

func (c *MenuController) Add() {
	c.jsonResult(consts.JRCodeSucc, "1", "b")
}

func (c *MenuController) Delete() {
	c.jsonResult(consts.JRCodeSucc, "1", "b")
}

func (c *MenuController) Edit() {
	c.jsonResult(consts.JRCodeSucc, "1", "b")
}
