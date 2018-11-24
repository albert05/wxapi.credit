package services

import (
	"github.com/astaxie/beego"
	"wxapi.credit/util/mysql"
	"encoding/gob"
	"wxapi.credit/models"
	_ "github.com/astaxie/beego/session/redis"
)

const LifeTime = 86400 * 29


func ConfigInit() {
	// init mysql
	dbDsn := beego.AppConfig.String("dbconfig::dsn")
	mysql.Init(dbDsn)

	InitSession()
}

func InitSession() {
	redisDsn := beego.AppConfig.String("redisconfig::dsn")

	gob.Register(models.User{})

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = LifeTime
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = LifeTime
	beego.BConfig.WebConfig.Session.SessionName = "gosessionid"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = redisDsn
}
