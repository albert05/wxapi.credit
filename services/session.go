package services

import (
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego"
)

var GS *session.Manager

func InitSession(dsn string) {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: dsn,
	}

	var err error
	GS, err = session.NewManager("redis", sessionConfig)
	if err != nil {
		beego.BeeLogger.Info(err.Error())
	}
	//go GS.GC()
}
