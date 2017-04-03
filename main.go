package main

import (
	_ "github.com/freewebsys/go-admin/routers"
	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"

	"github.com/astaxie/beego/orm"
)

func main() {

	//注册sqlite3
	orm.RegisterDataBase("default", "sqlite3", "go-admin.db")
	//同步 ORM 对象和数据库
	//这时, 在你重启应用的时候, beego 便会自动帮你创建数据库表。
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Run()
}
