package routers

import (
	"github.com/beego/beego/v2/server/web"

	"mall/controllers"
)

func init() {
	ns := web.NewNamespace("/mall",
		web.NSNamespace("/user",
			web.NSRouter("/register", &controllers.AccountController{}, "post:UserRegister"),
		),
	)
	web.AddNamespace(ns)
}
