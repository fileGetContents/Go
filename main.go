package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "quickstart/routers"
	"time"
)

func init() {
	//orm.Debug = true
	orm.RegisterDataBase("defult", "mysql", "root:@tcp(127.0.0.1:3306)/shop?charset=utf8mb4", 30)
	orm.RegisterDataBase("defult4", "mysql", "root:@tcp(127.0.0.1:3306)/shop?charset=utf8mb4", 30)
}

func main() {
	beego.Run()
	orm.DefaultTimeLoc = time.UTC
}
