package routers

import (
	"porject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{}, "get:Login")
	beego.Router("/user/register", &controllers.UserController{}, "get:Register")
	beego.Router("/user/register", &controllers.UserController{}, "post:RegisterUser")

}
