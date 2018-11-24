package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"wxapi.credit/services/wx"
	"wxapi.credit/services"
	"wxapi.credit/models"
)

// Operations about Users
type UserController struct {
	beego.Controller
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
		u.Data["json"] = services.FailedRetEx("login failed", map[string]interface{}{
			"err": "invalid params",
		})
		u.ServeJSON()
		return
	}

	r, err := wx.Login(params["code"])
	if  err != nil {
		u.Data["json"] = services.FailedRetEx("login failed", map[string]interface{}{
			"err": err.Error(),
		})
		u.ServeJSON()
		return
	}

	// insert into user table
	if ok := services.InsertUser(r.OpenId); !ok {
		u.Data["json"] = services.FailedRet("data insert failed")
		u.ServeJSON()
		return
	}

	// set session
	u.SetSession(beego.AppConfig.String("SessionKeyX"), r.OpenId)
	u.CruSession.SessionRelease(u.Ctx.ResponseWriter)

	u.Data["json"] = services.SuccRet("login success")
	u.ServeJSON()
}

// @Title Search
// @Description User Search
// @Success 200 {int}
// @Failure 403
// @router /search [post]
func (u *UserController) Search() {
	u.Data["json"] = services.SuccRetEx("login success", map[string]interface{}{
		"openid": models.OpenID,
	})
	u.ServeJSON()
}
