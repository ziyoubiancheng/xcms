package controllers

import (
	"github.com/ziyoubiancheng/xcms/consts"
	"github.com/ziyoubiancheng/xcms/models"
)

type MenuController struct {
	BaseController
}

func (c *MenuController) Index() {
	models.MenuModel.List()
	c.jsonResult(consts.JRCodeSucc, "1", "b")
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
