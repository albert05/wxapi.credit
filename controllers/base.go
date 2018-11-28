package controllers

import (
	"github.com/astaxie/beego"
	"wxapi.credit/models"
	"encoding/json"
	"wxapi.credit/services"
	"strings"
	"wxapi.credit/common"
)

const SessionKeyX  = "WX.USER"

// Operations about Users
type BaseController struct {
	beego.Controller
	IsLogin bool //标识 用户是否登陆
	User    *models.User //登陆的用户
	Params	map[string]string
}

func (ctx *BaseController) Prepare() {
	json.Unmarshal(ctx.Ctx.Input.RequestBody, &ctx.Params)

	ctx.IsLogin = false
	if u, ok := ctx.GetSession(SessionKeyX).(string); ok {
		user, err := models.FindUser(u)
		if user.Id > 0 && err == nil {
			ctx.User = user
			ctx.IsLogin = true
		}
	}

	ctx.MustLogin()
}

func (ctx *BaseController) MustParams(keys ...string) {
	if !common.CheckParams(ctx.Params, keys) {
		ctx.Abort666("failed", map[string]interface{}{
			"err": "invalid params",
		})
	}
}

func (ctx *BaseController) MustLogin() {
	uri := ctx.Ctx.Input.URI()
	uri = strings.TrimRight(strings.TrimLeft(uri, "/"), "/")
	uriList := services.LFilter.GetNLoginList()
	if _, ok := uriList[uri]; !ok && !ctx.IsLogin {
		ctx.JsonLogin()
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
	ctx.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")

	var data map[string]interface{}
	if len(datas) > 0 {
		data = datas[0]
	}

	ctx.CustomAbort(200, map2Json(&Code{
		Code:   FAILED_CODE,
		Message:    msg,
		Data: data,
	}))
}

func map2Json(data interface{}) string {
	content, _ := json.Marshal(data)

	return string(content)
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
	Code int						`json:"code"`
	Message string 					`json:"message"`
	Data  	map[string]interface{} 	`json:"data"`
}



