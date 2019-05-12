package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	cache "github.com/patrickmn/go-cache"
	"github.com/ziyoubiancheng/xcms/consts"
	"github.com/ziyoubiancheng/xcms/models"
	"github.com/ziyoubiancheng/xcms/utils"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           *models.AdminModel
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

	var m models.MenuModel
	c.Data["Menu"] = m.Tree()
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

func (c *BaseController) listJsonResult(code consts.JsonResultCode, msg string, count int, obj interface{}) {
	r := &models.ListJsonResult{code, msg, obj, count}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) auth() {
	arr := strings.Split(c.Ctx.GetCookie("auth"), "|")
	c.userId = 0
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			var err error

			cheUser, found := utils.Cache.Get("uid" + strconv.Itoa(userId))
			user := &models.AdminModel{}
			if found && cheUser != nil { //从缓存取用户
				user = cheUser.(*models.AdminModel)
			} else {
				user, err = user.GetById(userId)
				utils.Cache.Set("uid"+strconv.Itoa(userId), user, cache.DefaultExpiration)
			}
			if err == nil && password == utils.Md5([]byte(user.Password+user.Salt)) {
				c.userId = user.Id

				c.loginName = user.LoginName
				c.userName = user.RealName
				c.user = user
				//c.AdminAuth()
			}

			isHasAuth := strings.Contains(c.allowUrl, c.controllerName+"/"+c.actionName)
			//不需要权限检查
			noAuth := "ajaxsave/ajaxdel/table/loginin/loginout/getnodes/start/show/ajaxapisave/index/group/public/env/code/apidetail"
			isNoAuth := strings.Contains(noAuth, c.actionName)
			if isHasAuth == false && isNoAuth == false {
				c.Ctx.WriteString("没有权限")
				c.jsonResult(consts.JRCodeFailed, "", 0)
				return
			}
		}
	}

	if c.userId == 0 && (c.controllerName != "login" && c.actionName != "loginin") {
		c.redirect(beego.URLFor("LoginController.LoginIn"))
	}
}

// 重定向
func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

// 重定向 去错误页
func (c *BaseController) pageError(msg string) {
	errorurl := c.URLFor("HomeController.Error") + "/" + msg
	c.Redirect(errorurl, 302)
	c.StopRun()
}

// 重定向 去登录页
func (c *BaseController) pageLogin() {
	url := c.URLFor("HomeController.Login")
	c.Redirect(url, 302)
	c.StopRun()
}
