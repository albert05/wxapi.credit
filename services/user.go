package services

import (
	"wxapi.credit/models"
	"wxapi.credit/util/mysql"
)

func InsertUser(openid string) *models.User {
	user, err := models.FindUser(openid)
	if err != nil {
		return nil
	}

	if user.Id > 0 {
		return user
	}

	if u := models.InsertUser(mysql.MapModel{
		"openid": openid,
	}); u != nil {
		return u
	}

	return nil
}
