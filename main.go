package main

import (
	"MQConfig/controllers"
	_ "MQConfig/routers"
	"github.com/astaxie/beego"
)

func main() {
	//关闭自动渲染
	beego.BConfig.WebConfig.AutoRender = false

	//注册错误控制器
	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()
}

