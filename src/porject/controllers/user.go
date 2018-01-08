package controllers

import (
	"html/template"
	"strconv"

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
	u.XSRFExpire = 7200
	u.Data["xsrfdata"] = template.HTML(u.XSRFFormHTML())
	u.TplName = "login.html"
}

func (u *UserController) Loginin() {
	//get user login message
	userName := u.GetString("username")
	passWord := u.GetString("password")

	if ok, checkMsg := UserRegx(userName, passWord); !ok {
		u.Ctx.WriteString(checkMsg)
		u.StopRun()
	}

	if userName == "mark" && passWord == "123" {
		u.Ctx.WriteString("Logint ok !")
		u.Ctx.WriteString("LoginName:" + userName + "LoginPwd:" + passWord)
	} else {
		u.Ctx.WriteString("Login password error")
	}

}

func (u *UserController) Register() {
	u.Data["title"] = "Register Mark"
	u.XSRFExpire = 7200
	u.Data["xsrfdata"] = template.HTML(u.XSRFFormHTML())
	u.TplName = "register.html"
}

type user struct {
	Username interface{} `form:"username"` //解析时 首字母必须大写
	Password interface{} `form:"password"`
	Age      int         `form:"age"`
}

func (u *UserController) RegisterUser() {
	//get user register message  example:1
	userName := u.GetString("username")
	pwd := u.GetString("password")
	//get int type
	age := u.Input().Get("age")
	intAge, err := strconv.Atoi(age) //转为int64
	if err != nil && intAge < 1 {
		u.Ctx.WriteString("Age error")
	}

	//get user register message  example:2
	userMsg := user{}
	if err := u.ParseForm(&userMsg); err != nil {
		u.Ctx.WriteString("Register message error!")
	}

	u.Ctx.WriteString("LoginName:" + userName + "LoginPwd:" + pwd)
	u.Ctx.WriteString("Register ok!")
}

func UserRegx(username string, password string) (bool, string) {

	if username == "" {
		return false, "The username can not be empty !"
	}

	if password == "" {
		return false, "The password can not be empty !"
	}

	return true, ""
}
