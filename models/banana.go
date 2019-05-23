package models

import "github.com/astaxie/beego/orm"

type Bananas struct {
	BananaId     int `orm:"pk"`
	BananaImages string
	BananaUrl    string
	BananaSort   int
}

func init() {
	orm.RegisterModel(new(Bananas))
}

func (b *Bananas) TableName() string {
	return "bananas"
}

// 获取首页轮播图
func (b *Bananas) BananaGet() (bool, []orm.Params) {
	var banana Bananas
	var maps []orm.Params
	_, err := orm.NewOrm().QueryTable(banana).OrderBy("-BananaSort").Limit(5, 0).Values(&maps)
	if err == nil {
		return true, maps
	}
	return false, maps
}
