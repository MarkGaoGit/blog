package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.Data["Website"] = "Mark"
	c.Data["Email"] = "gaokang@powerlong.com"

	userMsg := c.GetSession("user")
	user, ok := userMsg.(map[string]interface{})
	fmt.Println(user["username"])
	fmt.Println(ok)

	if userMsg == nil {
		c.Data["loginUrl"] = c.URLFor("UserController.Login")
		c.Data["login"] = "LOGIN"
	} else {
		c.Data["loginUrl"] = c.URLFor("UserController.LoginOut")
		c.Data["login"] = userMsg

	}

	c.TplName = "index.html"
}
