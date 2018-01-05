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
	u.Data["title"] = "Login Mark"
	u.TplName = "login.html"
}

func (u *UserController) Loginin() {
	u.Ctx.WriteString("Login ok!")
}

func (u *UserController) Register() {
	u.Data["title"] = "Register Mark"
	u.TplName = "register.html"
}

func (u *UserController) RegisterUser() {
	u.Ctx.WriteString("Register ok!")
}
