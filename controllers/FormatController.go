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

func (c *FormatController) Edit() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "format/footerjs_edit.html"
	c.setTpl()
}
func (c *FormatController) EditDo() {

}

func (c *FormatController) Add() {

}
func (c *FormatController) AddDo() {

}

func (c *FormatController) DeleteDo() {

}
