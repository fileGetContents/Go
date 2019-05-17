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
	if len(prv) >= 0 {
		//	orm.Debug = true
		userModel := new(models.Users)
		o := orm.NewOrm()
		qs := o.QueryTable(userModel)
		user := models.Users{}
		qs.Filter("UserId", prv).Limit(1).One(&user)
		prv, _ := strconv.Atoi(prv) // 转化
		fmt.Println(prv)
		fmt.Println(user.UserId)
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
	d, _ := valid.Valid(&valiData)
	if !d {
		var res helps.Statuss
		for _, err := range valid.Errors {
			//log.Println(err.Key, err.Message)
			res.Statuss = append(res.Statuss, helps.Status{Msg: err.Message});
		}
		res.CodeNum = 1
		b, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
		this.Ctx.WriteString(string(b))
		this.StopRun()
	}
	//this.Ctx.WriteString(helps.Suc(""))
	//this.ServeJSON()
}
