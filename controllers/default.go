package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"quickstart/models"
)

type MainController struct {
	beego.Controller
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/shop?charset=utf8mb4", 30)
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Test() {
	//	o := orm.NewOrm()
	users := new(models.Users)
	name := users.TableName()
	fmt.Println(name)
	c.TplName = "test.tpl"
}
