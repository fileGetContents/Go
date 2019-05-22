package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/helps"
)

type ApiController struct {
	beego.Controller
}

// @Title 图形验证码
// @Description get all the staticblock by key
// @Param   key     path    string  true        "The email for login"
// @Success 200 {object}
// @Failure 404 不存在该接口
// @Failure 500 系统错误
// @router /external/sms/:phone [post]
func (this *ApiController) Capatcha() {
	idKey, image := helps.CodeCaptchaCreate()
	this.SetSession("Capatcha", idKey) //  设置session
	//helps.Suc(image)
	var status helps.Status
	status.Msg = image
	this.Data["json"] = &status
	this.ServeJSON()
}

// @Title 验证图形验证码
// @Description get all the staticblock by key
// @Param   key     path    string  true        "The email for login"
// @Success 200 {object}
// @Failure 404 不存在该接口
// @Failure 500 系统错误
// @router /external/sms/:phone [post]
func (this *ApiController) ValCaptcha() {
	cap := this.GetSession("Capatcha")
	var status helps.Status
	if cap == nil {
		status.Code = 1000
	}
	input := this.GetString("code")

	if input == cap {
		status.Code = 0
	} else {
		status.Code = 1000
	}
	this.Data["json"] = &status
	this.ServeJSON()
}
