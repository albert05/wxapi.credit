package controllers

import (
	"github.com/astaxie/beego"
	"wxapi.credit/models"
)

const SessionKeyX  = "WX.USER"

// Operations about Users
type BaseController struct {
	beego.Controller
	IsLogin bool //标识 用户是否登陆
	User    models.User //登陆的用户
}

func (ctx *BaseController) Prepare() {
	ctx.IsLogin = false
	if u, ok := ctx.GetSession(SessionKeyX).(string); ok {
		user, err := models.FindUser(u)
		if user.Id > 0 && err == nil {
			ctx.User = *user
			ctx.IsLogin = true
		}
	}
}


func (ctx *BaseController) JsonSucc(msg string, datas ...map[string]interface{}) {
	var data map[string]interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	ctx.Data["json"] = &Code{
		Code:   SUCCESS_CODE,
		Message:    msg,
		Data: data,
	}
	ctx.ServeJSON()
}

func (ctx *BaseController) Abort666(msg string, datas ...map[string]interface{}) {
	var data map[string]interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	ctx.Data["json"] = &Code{
		Code:   FAILED_CODE,
		Message:    msg,
		Data: data,
	}
	ctx.Abort("666")
}

func (ctx *BaseController) JsonLogin() {
	ctx.Data["json"] = &Code{
		Code:   LOGIN_CODE,
		Message:    "to login",
	}
	ctx.ServeJSON()
}

const SUCCESS_CODE  = 200
const FAILED_CODE  = 100
const LOGIN_CODE  = 300

type Code struct {
	Code int
	Message string
	Data  	map[string]interface{}
}



