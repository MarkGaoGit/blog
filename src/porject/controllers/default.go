package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.Data["Website"] = "Mark"
	c.Data["Email"] = "gaokang@powerlong.com"

	userMsg := c.GetSession("user")
	if userMsg == nil {
		c.Data["loginUrl"] = "/user"
		c.Data["login"] = "LOGIN"
	} else {
		c.Data["loginUrl"] = "/user/loginOut"
		c.Data["login"] = "LOGIN OUT"
	}

	c.TplName = "index.html"
}
