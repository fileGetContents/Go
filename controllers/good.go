package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/helps"
	"quickstart/models"
)

type GoodController struct {
	beego.Controller
}

func (this *GoodController) Details() {
	id := this.Ctx.Input.Param(":id")
	var good models.Goods
	b, data := good.GoodGetFindById(id)
	if b == false {
		this.Redirect("/", 302)
		return
	}
	ok, res := helps.JsonToSlice(helps.NilToString(data[0]["GoodImages"]))
	if ok {
		this.Data["images"] = res
	} else {
		this.Data["images"] = nil
	}


	this.Data["data"] = data[0]
	this.TplName = "goodDetails.html"
}
