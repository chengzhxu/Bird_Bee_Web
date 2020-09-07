package routers

import (
	"Bird_Bee_Web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/home", &controllers.UserController{}, "*:Home")

	beego.Include(&controllers.UserController{})
}
