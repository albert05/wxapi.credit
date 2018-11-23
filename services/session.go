package services

import "github.com/astaxie/beego/session"

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

	GS, _ = session.NewManager("mysql", sessionConfig)
	go GS.GC()
}
