package routers

import (
	"github.com/lakexyde/gosocial/controllers"
	"github.com/astaxie/beego"
	"github.com/lakexyde/gosocial/filters"
)

func init() {

	beego.InsertFilter("/:user/*", beego.BeforeRouter, filters.FilterUser)
	
	beego.Router("/", &controllers.MainController{}, "*:Homepage")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/register", &controllers.MainController{}, "post:Register")

    beego.Router("/:user/timeline", &controllers.UserController{}, "*:Timeline")
	beego.Router("/:user/profile", &controllers.UserController{}, "*:Profile")
}
