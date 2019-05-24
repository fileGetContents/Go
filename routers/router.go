package routers

import (
	"github.com/astaxie/beego"
	"quickstart/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "*:Home")

	beego.Router("/user", &controllers.MainController{}, "*:Test")

	beego.Router("/user/register/?:id:int", &controllers.RegisterController{}) // 用户注册

	beego.Router("/user/login", &controllers.UserController{}, "get:LoginGet")       // 用户登录
	beego.Router("/user/loginpost", &controllers.UserController{}, "post:LoginPost") // 用户登录

	beego.Router("/external/sms/:phone:string", &controllers.ExternalController{}, "*:Sms") // 短信验证码

	beego.Router("/api/capatcha", &controllers.ApiController{}, "*:Capatcha")     // 图形验证码
	beego.Router("/api/valcaptcha", &controllers.ApiController{}, "*:ValCaptcha") // 验证图形验证

	beego.Router("/good/details/:id:int", &controllers.GoodController{}, "get:Details") //
	beego.Router("/good/list", &controllers.GoodController{}, "get:List")
	beego.Router("/good/list/post", &controllers.GoodController{}, "*:ListPost")

}
