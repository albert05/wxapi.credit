package sessions

import (
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego"
	"wxapi.credit/models"
)

var GS *session.Manager
var OpenID string

// wx user sessions valid 30 days
const LifeTime = 86400 * 29
const SESSION_KEY = "WX.USER"

func InitSession(dsn string) {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:LifeTime,
		Maxlifetime: LifeTime,
		Secure: false,
		CookieLifeTime: LifeTime,
		ProviderConfig: dsn,
	}

	var err error
	GS, err = session.NewManager("redis", sessionConfig)
	if err != nil {
		beego.BeeLogger.Info(err.Error())
	}
	//go GS.GC()
}


// set user login session
func SetUser(s session.Store, openid string) {
	s.Set(SESSION_KEY, openid)
}

// check user is login
func IsLogin(s session.Store) bool {
	openid := s.Get(SESSION_KEY)

	if o, ok := openid.(string); ok {
		if u, err := models.FindUser(o); err == nil && u.Openid != "" {
			OpenID = o
			return true
		}
	}

	return false
}
