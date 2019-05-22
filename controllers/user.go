package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/helps"
	"quickstart/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) LoginGet() {
	this.TplName = "login.html"
}

// @Title 用户登录提交
// @Param phone 手机号
// @Param userpwd 密码
func (this *UserController) LoginPost() {
	phone := this.GetString("phone")
	pwd := this.GetString("userpwd")
	var UserModel models.Users
	b, user := UserModel.UserLogin(phone, pwd)
	var status helps.Status
	if b {
		status.Code = 0
		status.Msg = ""
		this.SetSession("user_id", user[0]["UserId"])
		this.SetSession("user_name", user[0]["UserName"])
		// other
	} else {
		status.Code = 1000
		status.Msg = "账号或密码错误"
	}
	this.Data["json"] = status
	this.ServeJSON()
}
