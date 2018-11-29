package controllers

import (
	"wxapi.credit/util/mysql"
	"wxapi.credit/models"
)

// Operations about Upload
type UploadController struct {
	BaseController
}

// @Title UploadUserInfo
// @Description User UploadUserInfo
// @Param formid query string true
// @Success 200 {int}
// @Failure 403
// @router /upload-formid [post]
func (u *UploadController) UploadFormid() {
	u.MustParams("formid")

	f, err := models.FindFormidInfo(u.Params["formid"])
	if f == nil && err == nil {
		models.InsertFormidInfo(mysql.MapModel{
			"formid": u.Params["formid"],
			"user_id": u.User.Id,
		})
	}

	u.JsonSucc("upload success")
}

// @Title UploadFeedback
// @Description User UploadFeedback
// @Param formid query string true
// @Success 200 {int}
// @Failure 403
// @router /upload-feedback [post]
func (u *UploadController) UploadFeedback() {
	u.MustParams("phone", "content")

	models.InsertFeedback(mysql.MapModel{
		"phone": u.Params["phone"],
		"content": u.Params["content"],
		"user_id": u.User.Id,
	})

	u.JsonSucc("upload success")
}
