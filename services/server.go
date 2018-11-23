package services

import (
	"github.com/astaxie/beego"
	"wxapi.credit/util/mysql"
	"wxapi.credit/services/sessions"
)

func ConfigInit() {
	// init mysql
	dbDsn := beego.AppConfig.String("dbconfig::dsn")
	mysql.Init(dbDsn)

	// init sessions
	redisDsn := beego.AppConfig.String("redisconfig::dsn")
	sessions.InitSession(redisDsn)
}
