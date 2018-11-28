package controllers

import (
	"wxapi.credit/services/wx"
	"wxapi.credit/services"
	"wxapi.credit/util/mysql"
)

// Operations about Users
type UserController struct {
	BaseController
}

// @Title Login
// @Description User Login
// @Param userCode query string true
// @Success 200 {int}
// @Failure 403
// @router /login [post]
func (u *UserController) Login() {
	u.MustParams("userCode")

	r, err := wx.Login(u.Params["userCode"])
	if  err != nil {
		u.Abort666("login failed", map[string]interface{}{
			"err": err.Error(),
		})
	}
	//
	//r := &wx.LoginResp{
	//	OpenId: "123456",
	//}

	// insert into user table
	user := services.InsertUser(r.OpenId);
	if user == nil {
		u.Abort666("data insert failed")
	}

	// set session
	u.SetSession(SessionKeyX, r.OpenId)

	u.JsonSucc("login success", map[string]interface{}{
		"uid": user.Id,
		"sessionid": u.CruSession.SessionID(),
		"remindStatus": 1, //TODO
	})
}

// @Title UploadUserInfo
// @Description User UploadUserInfo
// @Param nickName query string true
// @Param avatarUrl query string true
// @Param gender query string true
// @Param city query string true
// @Param province query string true
// @Param country query string true
// @Param language query string true
// @Success 200 {int}
// @Failure 403
// @router /upload-user-info [post]
func (u *UserController) UploadUserInfo() {
	u.MustParams("nickName", "avatarUrl", "gender", "city", "province", "country", "language")

	u.User.Update(mysql.MapModel{
		"nick_name": u.Params["nickName"],
		"avatar_url": u.Params["avatarUrl"],
		"gender": u.Params["gender"],
		"province": u.Params["city"],
		"city": u.Params["province"],
		"country": u.Params["country"],
		"language": u.Params["language"],
	})

	u.JsonSucc("upload success")
}

// @Title Remind
// @Description User Remind
// @Param status query int true
// @Success 200 {int}
// @Failure 403
// @router /remind [post]
func (u *UserController) Remind() {
	u.MustParams("status")

	u.User.Update(mysql.MapModel{
		"remind": u.Params["status"],
	})

	u.JsonSucc("upload success")
}

