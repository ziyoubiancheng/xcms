package controllers

import (
	"fmt"
	"strconv"

	//"github.com/astaxie/beego/orm"
	//"github.com/bitly/go-simplejson"
	"github.com/ziyoubiancheng/xcms/consts"
	"github.com/ziyoubiancheng/xcms/models"
)

type DataController struct {
	BaseController
	Mid int
}

func (c *DataController) Prepare() {
	c.BaseController.Prepare()

	midstr := c.Ctx.Input.Param(":mid")
	mid, err := strconv.Atoi(midstr)
	c.Data["Mid"] = midstr
	if nil != err || mid <= 0 {
		//TODO error page
		c.setTpl()
	}
	c.Mid = mid
}

func (c *DataController) Index() {
	sj := models.MenuFormatStruct(c.Mid)
	//fmt.Println(sj.Get("schema"))
	if sj != nil {
		title := make(map[string]string)
		titlemap := sj.Get("schema")
		for k, _ := range titlemap.MustMap() {
			stype := titlemap.GetPath(k, "type").MustString()
			fmt.Println(k)
			fmt.Println(stype)
			if "object" != stype && "array" != stype {
				title[k] = titlemap.GetPath(k, "title").MustString()
			}
		}

		fmt.Println(title)
		c.Data["Title"] = title
	}

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "data/footerjs.html"
	c.setTpl()

}

func (c *DataController) List() {

	c.listJsonResult(consts.JRCodeFailed, "nil", 0, nil)
}

func (c *DataController) Add() {
	format := models.MenuFormatStruct(c.Mid)
	c.Data["Schema"] = format.Get("schema").MustMap()
	c.Data["Form"] = format.Get("form").MustArray()

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "data/footerjs_add.html"
	c.setTpl("data/add.html", "common/layout_jfedit.html")
}
func (c *DataController) AddDo() {
	fmt.Println(c.Mid)
	fmt.Println("+++++++++")
	//fmt.Println(c.Ctx.Request)
	var m interface{}
	if err := c.ParseForm(&m); err == nil {
		fmt.Println(m)
	}
	fmt.Println("---------")
	c.listJsonResult(consts.JRCodeFailed, "nil", 0, nil)
}

func (c *DataController) Edit() {

}
func (c *DataController) EditDo() {

}

func (c *DataController) DeleteDo() {

}
