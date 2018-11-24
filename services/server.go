package services

import (
	"github.com/astaxie/beego"
	"wxapi.credit/util/mysql"
)

func ConfigInit() {
	// init mysql
	dbDsn := beego.AppConfig.String("dbconfig::dsn")
	mysql.Init(dbDsn)
}
