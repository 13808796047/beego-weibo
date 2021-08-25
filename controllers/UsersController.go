package controllers

import "github.com/astaxie/beego"

type UsersController struct {
	beego.Controller
}

func (c *UsersController) Signup() {
	c.Data["Title"] = "注册"
	c.Layout = "layouts/default.html"
	c.TplName = "users/create.html"
}
