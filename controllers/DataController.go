package controllers

import (
	//	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/bitly/go-simplejson"
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
			if "object" != stype && "array" != stype {
				title[k] = titlemap.GetPath(k, "title").MustString()
			}
		}
		c.Data["Title"] = title
	}

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "data/footerjs.html"
	c.setTpl()

}

func (c *DataController) List() {
	page, err := c.GetInt("page")
	if err != nil {
		page = 1
	}
	size, err := c.GetInt("limit")
	if err != nil {
		size = 20
	}

	data, total := models.DataList(c.Mid, size, page)
	c.listJsonResult(consts.JRCodeSucc, "ok", total, data)
}

func (c *DataController) Add() {
	c.initForm()

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "data/footerjs_add.html"
	c.setTpl("data/add.html", "common/layout_jfedit.html")
}
func (c *DataController) AddDo() {
	if len(c.Ctx.Input.RequestBody) > 0 {
		//map
		sj, err := simplejson.NewJson(c.Ctx.Input.RequestBody)
		if nil == err {
			//data model
			var m models.DataModel
			m.Content = string(c.Ctx.Input.RequestBody)
			//Form默认数据
			m.Mid = c.Mid
			m.Parent = sj.Get("parent").MustInt()
			m.Name = sj.Get("name").MustString()
			m.Seq = sj.Get("seq").MustInt()
			m.Status = int8(sj.Get("status").MustInt())
			m.UpdateTime = time.Now().Unix()

			o := orm.NewOrm()
			id, err := o.Insert(&m)
			if nil == err {
				c.jsonResult(consts.JRCodeSucc, "ok", id)
			}
		}
	}
	c.jsonResult(consts.JRCodeFailed, "nil", 0)
}

func (c *DataController) Edit() {
	c.initForm()

}
func (c *DataController) EditDo() {

}

func (c *DataController) DeleteDo() {

}

func (c *DataController) initForm() {
	format := models.MenuFormatStruct(c.Mid)
	if nil == format {
		return
	}
	schemaMap := format.Get("schema")
	formArray := format.Get("form")

	//添加通用Form
	fa := formArray.MustArray()
	if len(fa) <= 0 {
		var tmpArray []map[string]string
		tmpArray = append(tmpArray, map[string]string{"key": "parent"})
		tmpArray = append(tmpArray, map[string]string{"key": "name"})
		tmpArray = append(tmpArray, map[string]string{"key": "seq"})
		tmpArray = append(tmpArray, map[string]string{"key": "status"})
		for k, _ := range schemaMap.MustMap() {
			tmpArray = append(tmpArray, map[string]string{"key": k})
		}
		tmpArray = append(tmpArray, map[string]string{"type": "submit", "title": "提交"})

		c.Data["Form"] = tmpArray
	} else {
		var tmpArray []interface{}
		tmpArray = append(tmpArray, map[string]string{"key": "parent"})
		tmpArray = append(tmpArray, map[string]string{"key": "name"})
		tmpArray = append(tmpArray, map[string]string{"key": "seq"})
		tmpArray = append(tmpArray, map[string]string{"key": "status"})
		var haveSubmit bool = false
		for k, v := range formArray.MustArray() {
			tmpArray = append(tmpArray, v)
			tp := formArray.GetIndex(k).Get("type")
			if "submit" == tp.MustString() {
				haveSubmit = true
			}
		}
		if false == haveSubmit {
			tmpArray = append(tmpArray, map[string]string{"type": "submit", "title": "提交"})
		}
		c.Data["Form"] = tmpArray
	}

	//添加通用Schema
	schemaMap.SetPath([]string{"parent", "type"}, "integer")
	schemaMap.SetPath([]string{"parent", "title"}, "上级数据")

	schemaMap.SetPath([]string{"name", "type"}, "string")
	schemaMap.SetPath([]string{"name", "title"}, "名称")

	schemaMap.SetPath([]string{"seq", "type"}, "integer")
	schemaMap.SetPath([]string{"seq", "title"}, "排序(倒序)")

	schemaMap.SetPath([]string{"status", "type"}, "integer")
	schemaMap.SetPath([]string{"status", "title"}, "状态")
	schemaMap.SetPath([]string{"status", "enum"}, []int{0, 1})

	c.Data["Schema"] = schemaMap.MustMap()
}
