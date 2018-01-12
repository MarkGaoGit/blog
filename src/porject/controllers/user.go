package controllers

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
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
	flash := beego.ReadFromRequest(&u.Controller) //login error message
	u.Data["error"] = flash.Data["error"]
	u.TplName = "register.html"
}

type user struct {
	Username string `valid:"Required;Match(/\s/)"` //first not admin
	Password string `valid:"Required;Match(/\d\s/)"`
	Age      int    `valid:""` // age: 1 ~ 130
}

/*
 *	user register
 */
func (u *UserController) RegisterUser() {
	//get user register message  example:1
	userName := u.GetString("username")
	pwd := u.GetString("password")

	//get user register message  example:2
	userMsg := user{}
	if err := u.ParseForm(&userMsg); err != nil {
		u.Ctx.WriteString("Register message error!")
	}

	valid := validation.Validation{}
	res, err := valid.Valid(&userMsg)
	if err != nil {
		fmt.Println(err)
	}
	if !res {
		flash := beego.NewFlash() //flash 传递错误信息

		for _, err := range valid.Errors {
			fmt.Println(err.Key, err.Message) //print error
			flash.Error(err.Key + ":" + err.Message)
			flash.Store(&u.Controller)
		}
		u.Redirect("/user/register", 302)
	}

	userSession := make(map[string]string)
	userSession["username"] = userName
	userSession["password"] = pwd

	u.SetSession("user", userSession) //存入session

	u.Redirect(u.URLFor("MainController.Index"), 302)
}

func (u *user) Valid(v *validation.Validation) {
	if strings.Index(u.Username, "admin") != -1 {
		v.SetError("Username", "The username can not start with admin")
	}
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
