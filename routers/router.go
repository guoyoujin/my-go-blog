package routers

import (
	"my-go-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "GET:Index")
	beego.Router("/about", &controllers.IndexController{}, "GET:About")
}
