package filter

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"wxapi.credit/models"
	_ "github.com/astaxie/beego/session/redis"
)

// wx user sessions valid 30 days
const LifeTime = 86400 * 29

// not check login
var sFilter map[string]bool

func init() {
	sFilter = make(map[string]bool)
}

func Set(method string) {
	sFilter[method] = true
}

func SessionFilter(ctx *context.Context) {
	InitSession()
	ss, _ := GS.SessionStart(ctx.ResponseWriter, ctx.Request)

	if _, ok := sFilter[ctx.Input.URI()]; ok {
		return
	}

	key := beego.AppConfig.String("SessionKeyX")
	o, ok := ss.Get(key).(string)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/tologin")
	}

	models.OpenID = o
}

var GS *session.Manager


func InitSession() {
	redisDsn := beego.AppConfig.String("redisconfig::dsn")

	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:LifeTime,
		Maxlifetime: LifeTime,
		Secure: false,
		CookieLifeTime: LifeTime,
		ProviderConfig: redisDsn,
	}

	var err error
	GS, err = session.NewManager("redis", sessionConfig)
	if err != nil {
		beego.BeeLogger.Info(err.Error())
	}
	//go GS.GC()
}
