package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"quickstart/helps"
	"reflect"
)

type Goods struct {
	GoodId             int     `orm:"pk",json:"good_id"`
	GoodTitle          string  `json:"good_title"`
	GoodImage          string  `json:"good_image"`
	GoodImages         string  `json:"good_images"`
	GoodPrice          float64 `json:"good_price"`
	GoodMarketPrice    float64 `json:"good_market_price"`
	GoodText           string  `json:"good_text"`
	GoodStatus         int     `json:"good_status"`
	GoodShipping       int     `json:"good_shipping"`
	GoodFare           int     `json:"good_fare"`
	GoodTime           int     `json:"good_time"`
	GoodSore           int     `json:"good_sore"`
	GoodPayWay         int     `json:"good_pay_way"`
	GoodIntegral       int     `json:"good_integral"`
	GoodIntegralPrice  float64 `json:"good_integral_price"`
	GoodStock          int     `json:"good_stock"`
	GoodLimit          int     `json:"good_limit"`
	GoodReturnIntegral int     `json:"good_return_integral"`
	GoodMoneyOne       float64 `json:"good_money_one"`
	GoodMoneyTwo       float64 `json:"good_money_two"`
	GoodActivation     int     `json:"good_activation"`
	GoodRI             int     `json:"good_r_i"`
	GoodMO             float64 `json:"good_m_o"`
	GoodMT             float64 `json:"good_m_t"`
}

// json 数据结构
type GoodsJson struct {
	Code int     `json:"code"`
	Data []Goods `json:"data"`
}
// 输出数据转化器
func (g *Goods) GoodReturnJson(data []orm.Params) GoodsJson {
	var goodJson GoodsJson
	for _, v := range data {
		fmt.Println( v["GoodId"])

		reflect.TypeOf(v["GoodId"])

		goodJson.Data = append(goodJson.Data, Goods{
			GoodId:             helps.AssertionToInt(v["GoodId"]),
			GoodTitle:          helps.NilToString(v["GoodTitle"]),
			GoodImage:          helps.NilToString(v["GoodImage"]),
			GoodImages:         helps.NilToString(v["GoodImages"]),
			GoodPrice:          helps.AssertionToFloat64(v["GoodPrice"]),
			GoodMarketPrice:    helps.AssertionToFloat64(v["GoodMarketPrice"]),
			GoodText:           helps.NilToString(v["GoodText"]),
			GoodStatus:         helps.AssertionToInt(v["GoodStatus"]),
			GoodShipping:       helps.AssertionToInt(v["GoodShipping"]),
			GoodFare:           helps.AssertionToInt(v["GoodFare"]),
			GoodTime:           helps.AssertionToInt(v["GoodTime"]),
			GoodSore:           helps.AssertionToInt(v["GoodSore"]),
			GoodPayWay:         helps.AssertionToInt(v["GoodPayWay"]),
			GoodIntegral:       helps.AssertionToInt(v["GoodIntegral"]),
			GoodIntegralPrice:  helps.AssertionToFloat64(v["GoodIntegralPrice"]),
			GoodStock:          helps.AssertionToInt(v["GoodStock"]),
			GoodLimit:          helps.AssertionToInt(v["GoodLimit"]),
			GoodReturnIntegral: helps.AssertionToInt(v["GoodReturnIntegral"]),
			GoodMoneyOne:       helps.AssertionToFloat64(v["GoodMoneyOne"]),
			GoodMoneyTwo:       helps.AssertionToFloat64(v["GoodMoneyTwo"]),
			GoodActivation:     helps.AssertionToInt(v["GoodActivation"]),
			GoodRI:             helps.AssertionToInt(v["GoodRI"]),
			GoodMO:             helps.AssertionToFloat64(v["GoodMO"]),
			GoodMT:             helps.AssertionToFloat64(v["GoodMT"]),
		})
	}
	return goodJson
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
	_, err := orm.NewOrm().QueryTable(Good).Filter("GoodId", id).Limit(1, 0).Values(&maps)
	if err == nil && len(maps) != 0 {
		return true, maps
	}
	return false, maps
}

// 分页数字
func (g *Goods) GoodPage(Page, Limit int) []orm.Params {
	var Good Goods
	var maps []orm.Params
	if Page < 0 || Page == 1 || Page == 0 {
		Page = 0
	} else {
		Page = Page * Limit
	}
	orm.NewOrm().QueryTable(Good).Limit(Limit).Values(&maps)
	return maps
}
