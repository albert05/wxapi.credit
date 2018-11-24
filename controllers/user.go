package controllers

import (
	"encoding/json"
	"wxapi.credit/services/wx"
	"wxapi.credit/services"
	"wxapi.credit/models"
)

// Operations about Users
type UserController struct {
	BaseController
}

// @Title Login
// @Description User Login
// @Param code query string true
// @Success 200 {int}
// @Failure 403
// @router /login [post]
func (u *UserController) Login() {
	var params map[string]string
	json.Unmarshal(u.Ctx.Input.RequestBody, &params)

	if _, ok := params["code"]; !ok {
		u.Abort666("login failed", map[string]interface{}{
			"err": "invalid params",
		})
	}

	r, err := wx.Login(params["code"])
	if  err != nil {
		u.Abort666("data insert failed", map[string]interface{}{
			"err": err.Error(),
		})
	}

	// insert into user table
	if ok := services.InsertUser(r.OpenId); !ok {
		u.Abort666("data insert failed")
	}

	// set session
	u.SetSession(SessionKeyX, r.OpenId)

	u.JsonSucc("login success")
}

// @Title Search
// @Description User Search
// @Success 200 {int}
// @Failure 403
// @router /search [post]
func (u *UserController) Search() {
	if !u.IsLogin {
		u.JsonLogin()
	}

	u.JsonSucc("login success", map[string]interface{}{
		"openid": models.OpenID,
	})
}
