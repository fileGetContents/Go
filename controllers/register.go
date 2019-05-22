package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"quickstart/helps"
	"quickstart/models"
	"quickstart/validations"
	"strconv"
	"strings"
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
	prv, err := this.GetInt("prv", 0)
	if err != nil {
		prv = 0
	}
	//fmt.Println(data.UserId)
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
	json := &helps.Status{}
	////图形验证码
	cap := this.GetSession("Capatcha")
	if helps.VerfiyCaptcha(helps.NilToString(cap), smsimg) == false {
		json.Code = 1000
		json.Msg = "图形验证码错误"
		this.Data["json"] = json
		this.ServeJSON()
	}

	// 短信验证码
	code := helps.AssertionToInt(this.GetSession(phone))
	if !strings.EqualFold(strconv.Itoa(code), smscode) {
		json.Code = 1000
		json.Msg = "短信验证失败"
		this.Data["json"] = json
		this.ServeJSON()
	}

	//查询数据库
	var UsersModel models.Users
	b, _ := UsersModel.UserFindByUserPhone(phone)
	if !b {
		var user models.Users;
		user.UserName = ""
		user.UserPhone = phone
		user.UserPassword = newpwd
		user.UserSuperior = prv
		bo, _ := UsersModel.UserAddUser(&user)
		if bo {
			json.Code = 0
			json.Msg = ""
			this.Data["json"] = json
			this.ServeJSON()
		} else {
			json.Code = 1001
			json.Msg = "注册失败"
			this.Data["json"] = json
			this.ServeJSON()
		}
	} else {
		json.Code = 1000
		json.Msg = "已存在该账号"
		this.Data["json"] = json
		this.ServeJSON()
	}

}
