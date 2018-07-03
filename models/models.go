package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql",
		"root:root@tcp(127.0.0.1:3306)/godemo?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(Category), new(Topic))

	// create table
	orm.RunSyncdb("default", false, true)
}
