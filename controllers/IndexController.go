package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

//首页
func (c *IndexController) Index() {
	c.Data["PageTitle"] = "首页"
	c.Layout = "layout/layout.html"
	c.TplName = "index.html"
}

//关于
func (c *IndexController) About() {
    c.Data["PageTitle"] = "关于"
    c.Layout = "layout/layout.html"
    c.TplName = "about/about.html"
}
