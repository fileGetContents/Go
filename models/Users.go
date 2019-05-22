package models

import (
	"github.com/astaxie/beego/orm"
)

type Users struct {
	UserId     int `orm:"pk"`
	UserStatus bool
	UserName   string
	UserAvatar string
	UserPhone  string
	//DeletedAt    string
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

// @Title 获取一条用户记录
// @Param UserPhone 手机号
func (u *Users) UserFindByUserPhone(UserPhone string) (bool, []orm.Params) {
	var Users Users
	var maps []orm.Params
	_, err := orm.NewOrm().QueryTable(Users).Filter("UserPhone", UserPhone).Limit(1, 0).Values(&maps)
	if err != nil || len(maps) == 0 {
		return false, maps
	}
	return true, maps
}

// @Title 用户登录
// @Param  phone 手机号
// @Param  password 密码
func (u *Users) UserLogin(phone, password string) (bool, []orm.Params) {
	b, data := u.UserFindByUserPhone(phone)
	if b {
		if password == data[0]["UserPassword"] {
			return true, data
		} else {
			return false, data
		}
	} else {
		return false, data
	}
}

// @Title 添加一个用户
// @Param User 数据结构晶体
func (u *Users) UserAddUser(user interface{}) (bool, int64) {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	if err == nil {
		return true, id
	}
	return false, id
}
