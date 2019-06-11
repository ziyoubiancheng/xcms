package controllers

import (
	"strconv"

	//"github.com/astaxie/beego/orm"
	"github.com/ziyoubiancheng/xcms/models"
)

type DataController struct {
	BaseController
}

func (c *DataController) Index() {
	midstr := c.Ctx.Input.Param(":mid")
	mid, err := strconv.Atoi(midstr)
	if nil == err && mid > 0 {
		models.MenuStruct(mid)

		c.setTpl()
	} else {
		c.setTpl()
	}
}
