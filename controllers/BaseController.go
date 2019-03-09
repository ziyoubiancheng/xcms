package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ziyoubiancheng/xcms/models"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	curUser        models.UserModel
}
