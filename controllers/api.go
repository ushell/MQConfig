package controllers

import (
	"MQConfig/models"
	"bytes"
	"encoding/json"
	"time"
)

type ApiController struct {
	BaseController
}

type Command struct {
	AppId 	string `json:"app_id"`
	Command string `json:"command"`
	Key 	string `json:"key"`
	Value 	string `json:"value"`
}

func (this *ApiController) Get() {
	s := this.GetSession("x-request-token")

	if s == nil {
		this.SetSession("x-request-token", time.Now())
	}

	data := &Response{Code:0, Message:"获取数据成功", Data:s}
	this.ResponseJson(data)
}

func (this *ApiController) Post() {
	var command Command

	//参数
	body := this.Ctx.Input.RequestBody

	if bytes.Compare(body, nil) == 0 {
		data := &Response{Code:-1, Message:"参数为空"}
		this.ResponseJson(data)
	}

	if err := json.Unmarshal(body, &command); err != nil {
		data := &Response{Code:-1, Message:"参数格式异常"}
		this.ResponseJson(data)
	}

	data := &Response{Code:-1, Message:"获取数据失败", Data:command}
	this.ResponseJson(data)
}

func (this *ApiController) Index() {
	redis_data := _redis()

	data := &Response{Code:0, Message:"ok", Data:redis_data}
	this.ResponseJson(data)
}

func _redis() string {

	redisModel := new(models.RedisModel)

	data := redisModel.Demo3()

	return data
}