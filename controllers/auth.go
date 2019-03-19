package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"
)

type AuthController struct {
	BaseController
}

type AuthField struct {
	Username	string `json:"username"`
	Password	string `json:"password"`
	Token		string `json:"token"`
}

//认证
func (this *AuthController) Auth() {
	var auth AuthField

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &auth)

	if err != nil {
		data := &Response{Code:-1, Message:"参数异常"}
		this.ResponseJson(data)
	}

	//AUTH Token认证
	if auth.Token != "" {
		if auth.Token != "63a9f0ea7bb98050796b649e85481845" {
			data := &Response{Code:-1, Message:"身份凭证错误"}
			this.ResponseJson(data)
		}

	} else {
		if auth.Password == "" || auth.Username == "" {
			data := &Response{Code:-1, Message:"参数缺失"}
			this.ResponseJson(data)
		}

		//账户密码认证
		if auth.Username != "root" || auth.Password != "63a9f0ea7bb98050796b649e85481845" {
			data := &Response{Code:-1, Message:"账户或密码错误"}
			this.ResponseJson(data)
		}
	}

	token := make(map[string]string)

	hash := md5.New()

	io.WriteString(hash, strconv.FormatInt(time.Now().Unix(), 10))

	token["token"] = fmt.Sprintf("%x", hash.Sum(nil))

	data := &Response{Code:0, Message:"认证成功", Data:token}
	this.ResponseJson(data)
}
