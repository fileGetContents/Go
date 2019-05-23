package models

import "github.com/astaxie/beego/orm"

type Goods struct {
	GoodId             string `orm:"pk"`
	GoodTitle          string
	GoodImage          string
	GoodImages         string
	GoodPrice          float64
	GoodMarketPrice    float64
	GoodText           string
	GoodStatus         int
	GoodShipping       int
	GoodFare           int
	GoodTime           int
	GoodSore           int
	GoodPayWay         int
	GoodIntegral       int
	GoodIntegralPrice  float64
	GoodStock          int
	GoodLimit          int
	GoodReturnIntegral int
	GoodMoneyOne       float64
	GoodMoneyTwo       float64
	GoodActivation     int
	GoodRI             int
	GoodMO             float64
	GoodMT             float64
}

func init() {
	orm.RegisterModel(new(Goods))

}

func (g *Goods) TableName() string {
	return "goods"
}

// 获取首页推荐信息
func (g *Goods) GoodGetHome() (bool, []orm.Params) {
	var Good Goods
	var maps []orm.Params
	_, err := orm.NewOrm().QueryTable(Good).Filter("GoodStatus", 1).OrderBy("-GoodSore").Limit(20, 0).Values(&maps)
	if err == nil {
		return true, maps
	}
	return false, maps
}

// 获取一条数据
func (g *Goods) GoodGetFindById(id string) (bool, []orm.Params) {
	var Good Goods
	var maps []orm.Params
	_, err := orm.NewOrm().QueryTable(Good).Filter("GoodId", id).Limit(1,0).Values(&maps)
	if err == nil && len(maps) != 0 {
		return true, maps
	}
	return false, maps
}
