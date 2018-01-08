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
	c.TplName = "index.html"
}
