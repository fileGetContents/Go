package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/helps"
	"quickstart/models"
)

type GoodController struct {
	beego.Controller
}


func (this *GoodController) List() {
	var good models.Goods
	data := good.GoodPage(1, 10)

	this.Data["good"] = data
	this.TplName = "GoodList.html"
}

func (this *GoodController) ListPost() {
	var good models.Goods
	_, Page := this.GetInt("page", 0)
	_, Limit := this.GetInt("limit", 10)
	data := good.GoodPage(helps.AssertionToInt(Page), helps.AssertionToInt(Limit))
	jsonReturn := good.GoodReturnJson(data)
	this.Data["json"] = jsonReturn
	this.ServeJSON()
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
