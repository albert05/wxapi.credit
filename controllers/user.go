package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"wxapi.credit/services/sessions"
	"github.com/astaxie/beego/session"
	"wxapi.credit/services/wx"
	"wxapi.credit/services"
)

// Operations about Users
type UserController struct {
	beego.Controller
	SS session.Store
}

func (u *UserController) Init() {
	s, _ := sessions.GS.SessionStart(u.Ctx.ResponseWriter, u.Ctx.Request)
	u.SS = s
}

// @Title Login
// @Description User Login
// @Param code query string true
// @Success 200 {int}
// @Failure 403
// @router /login [post]
func (u *UserController) Login() {
	defer u.SS.SessionRelease(u.Ctx.ResponseWriter)

	var params map[string]string
	json.Unmarshal(u.Ctx.Input.RequestBody, &params)

	if _, ok := params["code"]; !ok {
		u.Data["json"] = services.FailedRetEx("login failed", map[string]interface{}{
			"err": "invalid params",
		})
		u.ServeJSON()
	}

	r, err := wx.Login(params["code"])
	if  err != nil {
		u.Data["json"] = services.FailedRetEx("login failed", map[string]interface{}{
			"err": err.Error(),
		})
		u.ServeJSON()
	}

	// insert into user table
	if ok := services.InsertUser(r.OpenId); !ok {
		u.ServeJSON()
	}

	sessions.SetUser(u.SS, r.OpenId)

	u.Data["json"] = services.SuccRet("login success")
	u.ServeJSON()
}

// @Title Login
// @Description User Search
// @Success 200 {int}
// @Failure 403
// @router /search [post]
func (u *UserController) Search() {
	if !sessions.IsLogin(u.SS) {
		u.Data["json"] = services.FailedRet("need login")
		u.ServeJSON()
	}

	u.Data["json"] = services.SuccRetEx("login success", map[string]interface{}{
		"openid": sessions.OpenID,
	})
	u.ServeJSON()
}
