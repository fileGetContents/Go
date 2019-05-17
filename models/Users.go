package models

import "github.com/astaxie/beego/orm"

type Users struct {
	UserId       int `orm:"pk"`
	UserStatus   bool
	UserName     string
	UserAvatar   string
	UserPhone    string
	DeletedAt    string
	UserTime     int
	UserPassword string
	UserIntegral float64
	UserMoney    float64
	UserSuperior int
}

func init() {
	orm.RegisterModel(new(Users))
}

func (u *Users) TableName() string {
	return "users"
}
