package routers

import (
	"beego-weibo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{}, "get:home")
	beego.Router("/", &controllers.StaticPageController{}, "get:Home")
	beego.Router("/help", &controllers.StaticPageController{}, "get:Help")
	beego.Router("/about", &controllers.StaticPageController{}, "get:About")
	beego.Router("signup", &controllers.UsersController{}, "get:Signup")
}
