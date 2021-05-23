package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)


func InitOrm() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "cychust:cyc921507@/chat?charset=utf8")
	orm.RegisterModel(new(User))
	// create table
	orm.RunSyncdb("default", false, true)
	o := orm.NewOrm()

	configInfo := beego.AppConfig.String("MYSQL::dbname")
	o.Using(configInfo)
}
