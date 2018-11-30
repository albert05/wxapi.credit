package services

import (
	"github.com/astaxie/beego"
	"wxapi.credit/util/mysql"
	"encoding/gob"
	"wxapi.credit/models"
	_ "github.com/astaxie/beego/session/redis"
	"fmt"
	"wxapi.credit/util/redis"
	"wxapi.credit/common"
)

const LifeTime = 86400 * 29


func ConfigInit() {
	// init mysql
	dbDsn := beego.AppConfig.String("dbconfig::dsn")
	mysql.Init(dbDsn)

	host := beego.AppConfig.String("redisconfig::host")
	password := beego.AppConfig.String("redisconfig::password")
	pool := beego.AppConfig.String("redisconfig::pool")
	db := beego.AppConfig.String("redisconfig::db")

	// init redis
	redis.Init(host, password, common.Str2Int(db))

	// init session
	InitSession(fmt.Sprintf("%s,%s,%s", host, pool, password))
}

func InitSession(dsn string) {
	gob.Register(models.User{})

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = LifeTime
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = LifeTime
	beego.BConfig.WebConfig.Session.SessionName = "gosessionid"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = dsn
}
