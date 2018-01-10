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

/*
 *	user loginPage
 */
func (u *UserController) Login() {

	userMsg := u.GetSession("user")
	if userMsg == nil {
		u.Data["title"] = "Login Mark"
		u.XSRFExpire = 7200
		u.Data["xsrfdata"] = template.HTML(u.XSRFFormHTML())

		flash := beego.ReadFromRequest(&u.Controller) //login error message
		u.Data["error"] = flash.Data["error"]

		u.TplName = "login.html"
	} else {
		u.Data["user"] = userMsg
		u.TplName = "loginok.html"

	}

}

/*
 *	user login
 */
func (u *UserController) Loginin() {
	//get user login message
	userName := u.GetString("username")
	passWord := u.GetString("password")

	if ok, checkMsg := UserRegx(userName, passWord); !ok {
		flash := beego.NewFlash() //flash 传递错误信息
		flash.Error(checkMsg)
		flash.Store(&u.Controller)
		u.Redirect("/user", 302)

	}

	if userName == "mark" && passWord == "123" {

		userMsg := make(map[string]string)
		userMsg["username"] = userName
		userMsg["password"] = passWord

		u.SetSession("user", userMsg) //存入session

		u.Data["user"] = userMsg
		u.Data["loginMsg"] = "Login ok!"

	} else {
		u.Data["loginMsg"] = "Login password error"
	}
	u.TplName = "loginok.html"

}

/*
 *	user loginOut
 */
func (u *UserController) LoginOut() {
	u.DelSession("user")
	u.TplName = "login.html"
}

/*
 *	user registerPage
 */
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

/*
 *	user register
 */
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

/*
 *	check user register
 */
func UserRegx(username string, password string) (bool, string) {

	if username == "" {
		return false, "The username can not be empty !"
	}

	if password == "" {
		return false, "The password can not be empty !"
	}

	return true, ""
}
