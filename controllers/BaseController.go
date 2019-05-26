package controllers

import (
	"fmt"
	//	"strconv"
	"strings"

	"github.com/astaxie/beego"
	//	cache "github.com/patrickmn/go-cache"
	"github.com/ziyoubiancheng/xcms/consts"
	"github.com/ziyoubiancheng/xcms/models"
	//	"github.com/ziyoubiancheng/xcms/utils"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           *models.UserModel
	userId         int
	userName       string
	loginName      string
	pageSize       int
	allowUrl       string
}

func (c *BaseController) Prepare() {
	//附值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	beego.Informational(c.controllerName, c.actionName)
	// TODO 保存用户数据
	//c.auth()
	fmt.Println("beego:perpare:" + c.controllerName + "," + c.actionName)

	c.Data["Menu"] = models.MenuStruct()
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

func (c *BaseController) auth() {
	if c.userId == 0 && (c.controllerName != "login") {
		c.pageLogin()
	}
}

// 重定向 去错误页
func (c *BaseController) pageError(msg string) {
	errorurl := c.URLFor("HomeController.Error") + "/" + msg
	c.Redirect(errorurl, 302)
	c.StopRun()
}

// 重定向 去登录页
func (c *BaseController) pageLogin() {
	url := c.URLFor("LoginController.Login")
	c.Redirect(url, 302)
	c.StopRun()
}
