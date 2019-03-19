package routers

import (
	"MQConfig/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/api/auth", &controllers.AuthController{}, "*:Auth")

	//API V1
	ns := beego.NewNamespace("/api/v1",
		beego.NSRouter("/get", &controllers.ApiController{}, "*:Get"),
		beego.NSRouter("/post", &controllers.ApiController{}, "*:Post"),
		beego.NSRouter("/", &controllers.ApiController{}, "*:Index"),
	)

	//注册
	beego.AddNamespace(ns)
}
