package main

import (
	_ "my-go-blog/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"my-go-blog/models"
)

func init(){
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@/"+ beego.AppConfig.String("mysqldb") +"?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai", 30)
	orm.RegisterModel(
		new(models.Users),
		new(models.Roles),
		new(models.RolePermissions),
		new(models.Permissions))
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}


