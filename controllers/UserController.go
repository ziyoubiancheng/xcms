package controllers

import (
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

	var m models.UserModel
	result, count := m.List(size, page)
	c.listJsonResult(consts.JRCodeSucc, "ok", count, result)
}
