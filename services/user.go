package services

import (
	"wxapi.credit/models"
	"wxapi.credit/util/mysql"
)

func InsertUser(openid string) bool {
	user, err := models.FindUser(openid)
	if err != nil {
		return false
	}

	if user.Id > 0 {
		return true
	}

	if u := models.InsertUser(mysql.MapModel{
		"openid": openid,
	}); u == nil {
		return false
	}

	return true
}
