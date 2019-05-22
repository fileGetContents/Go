package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.MainController{}, "*:Test")

	beego.Router("/user/register/?:id:int", &controllers.RegisterController{}) // 用户注册

	beego.Router("/user/login", &controllers.UserController{}, "*:LoginGet")    // 用户登录
	beego.Router("/user/loginpost",&controllers.UserController{},"*:LoginPost")


	beego.Router("/external/sms/:phone:string", &controllers.ExternalController{}, "*:Sms") // 短信验证码

	beego.Router("/api/capatcha", &controllers.ApiController{}, "*:Capatcha")     // 图形验证码
	beego.Router("/api/valcaptcha", &controllers.ApiController{}, "*:ValCaptcha") // 验证图形验证码

}
