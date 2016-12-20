package main

import (
	_ "my-go-blog/routers"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"my-go-blog/models"
)

func init(){
	//orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@/"+ beego.AppConfig.String("mysqldb") +"?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai", 30)
	//orm.RegisterModel(
	//	new(models.Users),
	//	new(models.Roles),
	//	new(models.RolePermissions),
	//	new(models.Permissions))
	//orm.RunSyncdb("default", false, true)
}


func main() {
	db, err := gorm.Open("mysql", beego.AppConfig.String("mysqluser")+":"+ beego.AppConfig.String("mysqlpass")+"@/"+beego.AppConfig.String("mysqldb")+"?charset=utf8&parseTime=True&loc=Local")
	if (err != nil) {
		panic("failed to connect database"+err.Error())
	}

	db.AutoMigrate(&models.User{})
	// Add table suffix when create tables
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Users{})
	defer db.Close()
	beego.Run()
}


