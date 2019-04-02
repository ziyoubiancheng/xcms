package controllers

//import (
//	"fmt"

//	"github.com/astaxie/beego/orm"
//	"github.com/ziyoubiancheng/xcms/consts"
//	"github.com/ziyoubiancheng/xcms/models"
//)

type FormatController struct {
	BaseController
}

func (c *FormatController) Index() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "format/footerjs.html"
	c.setTpl()
}
