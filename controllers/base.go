package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//响应基础数据结构
type Response struct {
	Code int 			`json:"code"`
	Message string 		`json:"message"`
	Data interface{} 	`json:"data,omitempty"`
}

func (c *BaseController) Prepare() {
	c.auth()
}

//身份校验
func (c *BaseController) auth() {
	token := c.Ctx.Request.Header.Get("x-request-token")

	if token != "a9b2629c150b60fdaff36b1c148cdaad" {
		c.Data["json"] = &Response{Code:-1, Message:"身份验证失败"}
		c.ServeJSON()
	}

}

func (c *BaseController) ResponseJson(data interface{}) {
	c.Data["json"] = data
	c.ServeJSON()
}
