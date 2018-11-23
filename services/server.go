package services

import (
	"github.com/astaxie/beego"
	"wxapi.credit/util/mysql"
	_ "github.com/astaxie/beego/session/redis"
)

func ConfigInit() {
	// init mysql
	dbDsn := beego.AppConfig.String("dbconfig::dsn")
	mysql.Init(dbDsn)

	// init session

	redisDsn := beego.AppConfig.String("redisconfig::dsn")
	InitSession(redisDsn)
}
