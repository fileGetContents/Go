package controllers

import (
	"github.com/astaxie/beego"
	"quickstart/models"
)

type IndexController struct {
	beego.Controller
}

// @Title 主页
func (this *IndexController) Home() {
	var GoodsModel models.Goods
	_, data := GoodsModel.GoodGetHome()
	this.Data["goods"] = data
	var Bananas models.Bananas
	_, banana := Bananas.BananaGet()
	this.Data["bananas"] = banana
	this.TplName = "index.html"
}
