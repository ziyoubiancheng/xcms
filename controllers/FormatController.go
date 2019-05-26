package controllers

import (
	//	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/ziyoubiancheng/xcms/consts"
	"github.com/ziyoubiancheng/xcms/models"
)

type FormatController struct {
	BaseController
}

func (c *FormatController) Edit() {
	c.Data["Mid"] = c.GetString("mid")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "format/footerjs_edit.html"
	c.setTpl("format/edit.html", "common/layout_edit.html")
}
func (c *FormatController) EditDo() {
	mid, _ := c.GetInt("mid")
	f := c.GetString("formatstr")
	//	fmt.Println(mid)
	//	fmt.Println(f)
	if 0 != mid {
		m := models.MenuModel{Mid: mid, Format: f}
		mid, _ := orm.NewOrm().Update(&m, "format")
		c.jsonResult(consts.JRCodeSucc, "ok", mid)
	}

	c.jsonResult(consts.JRCodeFailed, "", 0)
}
