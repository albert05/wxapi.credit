package controllers

import (
	"wxapi.credit/models"
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	params := u.Ctx.Input.Params()
	user, err := models.FindUser(params["openid"])
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = user
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by openid
// @Param	openid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :openid is empty
// @router /:openid [get]
func (u *UserController) Get() {
	uid := u.GetString(":openid")
	if uid != "" {
		user, err := models.FindUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}
