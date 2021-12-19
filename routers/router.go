package routers

import (
	"github.com/beego/beego/v2/server/web"

	"mall/controllers"
)

func init() {
	ns := web.NewNamespace("/mall",
		web.NSNamespace("/user",
			web.NSRouter("/register", &controllers.PassportController{}, "post:UserRegister"),
			web.NSRouter("/captcha", &controllers.PassportController{}, "get:NewCaptcha"),
		),
	)
	web.AddNamespace(ns)
}
