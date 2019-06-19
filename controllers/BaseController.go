package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/ziyoubiancheng/xcms/consts"
	"github.com/ziyoubiancheng/xcms/models"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (c *BaseController) Prepare() {
	//附值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	beego.Informational(c.controllerName, c.actionName)
	user := c.auth()
	c.Data["Menu"] = models.MenuTreeStruct(user)
}

// 设置模板
// 第一个参数模板，第二个参数为layout
func (c *BaseController) setTpl(template ...string) {
	var tplName string
	layout := "common/layout.html"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		//不要"Controller"这个10个字母
		ctrlName := strings.ToLower(c.controllerName[0 : len(c.controllerName)-10])
		actionName := strings.ToLower(c.actionName)
		tplName = ctrlName + "/" + actionName + ".html"
	}

	_, found := c.Data["Footer"]
	if !found {
		c.Data["Footer"] = "menu/footerjs.html"
	}
	c.Layout = layout
	c.TplName = tplName
}
func (c *BaseController) jsonResult(code consts.JsonResultCode, msg string, obj interface{}) {
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) listJsonResult(code consts.JsonResultCode, msg string, count int64, obj interface{}) {
	r := &models.ListJsonResult{code, msg, count, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) auth() models.UserModel {
	user := c.GetSession("xcmsuser")
	beego.Debug("base auth" + c.controllerName)
	beego.Debug(user)
	if user == nil {
		c.Redirect("/login", 302)
		c.StopRun()
		return models.UserModel{}
	} else {
		return user.(models.UserModel)
	}

}
