package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	this.Data["json"] = &Response{Code:-1, Message:"接口不存在"}
	this.ServeJSON()
}

func (this *ErrorController) Error501() {
	this.Data["json"] = &Response{Code:-1, Message:"服务异常"}
	this.ServeJSON()
}
