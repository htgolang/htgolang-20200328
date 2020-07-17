package init

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func init() {
	beego.BConfig.WebConfig.Session.SessionOn = beego.AppConfig.DefaultBool("session::SessionOn", false)
	beego.BConfig.WebConfig.Session.SessionProvider = beego.AppConfig.DefaultString("session::ProvideName", "memory")
	beego.BConfig.WebConfig.Session.SessionName = beego.AppConfig.DefaultString("session::Name", "TickID")
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = beego.AppConfig.DefaultInt64("session::GCLifeTime", 3600)
	beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.DefaultString("session::ProviderConfig", "")
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = beego.AppConfig.DefaultInt("session::CookieLifeTime", 3600)
}
