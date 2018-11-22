package services

import (
	"github.com/astaxie/beego"
	"wxapi.credit/util/mysql"
)

func ConfigInit() {
	// init mysql
	dsn := beego.AppConfig.String("dbconfig::dsn")
	mysql.Init(dsn)

	// init session
	InitSession()
}
