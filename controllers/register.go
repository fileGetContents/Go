package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"quickstart/helps"
	"quickstart/models"
	"quickstart/validations"
	"strconv"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.Data["prv"] = 0
	prv := this.Ctx.Input.Param(":id")
	if len(prv) > 0 {
		//	orm.Debug = true
		userModel := new(models.Users)
		o := orm.NewOrm()
		qs := o.QueryTable(userModel)
		user := models.Users{}
		qs.Filter("UserId", prv).Limit(1).One(&user)
		prv, _ := strconv.Atoi(prv) // 转化
		if prv == user.UserId && user.UserId != 0 {
			this.Data["prv"] = user.UserId
		}
	}

	this.TplName = "register.html"
}

func (this *RegisterController) Post() {
	phone := this.GetString("phone")
	newpwd := this.GetString("newpwd")
	smsimg := this.GetString("smsimg")
	smscode := this.GetString("smscode")
	// 数据验证
	valiData := validations.Register{
		Phone: phone, Newpwd: newpwd, Smsimg: smsimg, Smscode: smscode,
	}
	valid := validation.Validation{}
	valid.Valid(&valiData)
	if valid.HasErrors() {
		var error validations.Errors
		validMsg := validations.RegisterMsg()
		error.Code = validations.DefaultCode()
		for _, err := range valid.Errors {
			error.Es = append(error.Es, validations.Error{Key: err.Key, Msg: validMsg[err.Key]})
		}
		res, _ := json.Marshal(error)
		this.Ctx.WriteString(string(res))
		this.StopRun()
	}
	// 短信验证码
	json := &helps.Status{}
	code := this.GetSession(phone)
	fmt.Println(code)
	fmt.Println(smscode)
	if code != smscode {
		json.Code = 10000
		json.Msg = SmsMsgError
	}

	this.Data["json"] = json
	this.ServeJSON()
	//this.StopRun()
}
