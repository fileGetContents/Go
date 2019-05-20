package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.MainController{}, "*:Test")

	beego.Router("/register/?:id:int", &controllers.RegisterController{})


	beego.Router("/external/sms/:phone:string", &controllers.ExternalController{}, "*:Sms")
}
