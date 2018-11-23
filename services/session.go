package services

import "github.com/astaxie/beego/session"

var GS *session.Manager

func InitSession() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}

	GS, _ = session.NewManager("file", sessionConfig)
	go GS.GC()
}
