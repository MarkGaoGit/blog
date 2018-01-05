package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	u.Data["msg"] = "Test project"
	u.TplName = "user.html"
	//也可以使用 u.Ctx.WriteString("XXX") 直接输出
}

func (u *UserController) Login() {
	u.Data["msg"] = "This is login page"
	u.TplName = "user.html"
}

func (u *UserController) Register() {
	u.Data["msg"] = "This is register page"
	u.TplName = "user.html"
}

func (u *UserController) RegisterUser() {
	u.Data["msg"] = "This is create user page"
	u.TplName = "user.html"
}
