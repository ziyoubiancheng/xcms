package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/ziyoubiancheng/xcms/models"
	"github.com/ziyoubiancheng/xcms/utils"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Index() {
	if c.Ctx.Request.Method == "POST" {
		userkey := strings.TrimSpace(c.GetString("userkey"))
		password := strings.TrimSpace(c.GetString("password"))

		if len(userkey) > 0 && len(password) > 0 {
			passmd5 := utils.Md5([]byte(password))
			user := models.GetUserByName(userkey)

			fmt.Println(passmd5 + "-" + user.Password)
			if passmd5 == user.Password {
				user.Password = ""
				c.SetSession("xcmsuser", user)
				fmt.Println("login ok")
				c.Redirect("/", 302)
				c.StopRun()
			}
		}
	}
	c.TplName = "login/index.html"
}
