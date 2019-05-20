package controllers

// 外部接口
import (
	"github.com/astaxie/beego"
	"quickstart/helps"
	"regexp"
	"strconv"
)

const PhoneRule = "^1[34578]\\d{9}$"
const SmsMsgError  = "短信验证码错误"


type ExternalController struct {
	beego.Controller
}

// @Title 发送短信验证码
// @Description get all the staticblock by key
// @Param   key     path    string  true        "The email for login"
// @Success 200 {object}
// @Failure 404 不存在该接口
// @Failure 500 系统错误
// @router /external/sms/:phone [post]
func (this *ExternalController) Sms() {
	Phone := this.Ctx.Input.Param(":phone")
	b, _ := regexp.MatchString(PhoneRule, Phone)
	json := &helps.Status{}
	if b == false {
		json.Code = 1000
		json.Msg = "手机号格式错误"
	} else {
		num := helps.Random(1000, 9999)
		this.SetSession(Phone, num) // 设置session
		json.Msg = strconv.Itoa(num)
	}
	this.Data["json"] = json
	this.ServeJSON()
}
