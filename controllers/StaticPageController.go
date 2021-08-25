package controllers

import "github.com/astaxie/beego"

type StaticPageController struct {
	beego.Controller
}

func (c *StaticPageController) Home() {
	c.Data["Title"] = "首页"
	c.Layout = "layouts/default.html"
	c.TplName = "static_pages/home.html"
}
func (c *StaticPageController) Help() {
	c.Data["Title"] = "用户管理"
	c.Layout = "layouts/default.html"
	c.TplName = "static_pages/help.tpl"
}
func (c *StaticPageController) About() {
	c.Data["Title"] = "关于我"
	c.Layout = "layouts/default.html"
	c.TplName = "static_pages/about.html"
}
