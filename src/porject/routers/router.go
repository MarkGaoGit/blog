package routers

import (
	"porject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/user", &controllers.UserController{}, "get:Login")
	beego.Router("/user/login", &controllers.UserController{}, "post:Loginin")
	beego.Router("/user/loginOut", &controllers.UserController{}, "*:LoginOut")
	beego.Router("/user/register", &controllers.UserController{}, "get:Register")
	beego.Router("/user/register", &controllers.UserController{}, "post:RegisterUser")

}
