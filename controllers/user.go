package controllers

import (
	"wxapi.credit/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"wxapi.credit/services"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param openid query string true
// @Success 200 {int}
// @Failure 403
// @router /test [post]
func (u *UserController) Test() {
	var params map[string]string
	json.Unmarshal(u.Ctx.Input.RequestBody, &params)
	u.Data["json"] = params
	u.ServeJSON()

	user, err := models.FindUser(params["openid"])
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = user
	}
	u.ServeJSON()
}

// @Title Login
// @Description User Login
// @Param openid query string true
// @Param password  query string true
// @Success 200 {int}
// @Failure 403
// @router /login [post]
func (u *UserController) Login() {
	var params map[string]string
	json.Unmarshal(u.Ctx.Input.RequestBody, &params)

	sess, _ := services.GS.SessionStart(u.Ctx.ResponseWriter, u.Ctx.Request)
	defer sess.SessionRelease(u.Ctx.ResponseWriter)

	sessionId := sess.Get(params["openid"])
	if sessionId == nil {
		sessionId = sess.SessionID()
		sess.Set(params["openid"], sessionId)
	}

	//user, err := models.FindUser(params["openid"])

	response := make(map[string]interface{})
	response["sessionid"] = sessionId
	u.Data["json"] = response

	u.ServeJSON()
}
