package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Index() {
	if c.Ctx.Request.Method == "POST" {
		username := strings.TrimSpace(c.GetString("username"))
		password := strings.TrimSpace(c.GetString("password"))

		fmt.Println(username + password)

	}
	c.TplName = "login/index.html"
}
