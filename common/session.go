package common

import (
	"echo_test/config"
	"github.com/labstack/echo"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
)

//初始化SESSION
func InitSession(e *echo.Echo) {
	session_cookie_name := config.Config.SessionCookieName
	session_path := config.Config.SessionPath
	//设置session中间件
	//这里使用的session中间件，session数据保存在指定的目录
	e.Use(session.Middleware(sessions.NewFilesystemStore(session_path, []byte(session_cookie_name))))
}

//获取
func GetSession(ctx echo.Context, key string) interface{} {
	sess, _ := session.Get(key, ctx)
	return sess.Values[key]
}

//设置
func SetSession(ctx echo.Context, key string, value interface{}) {
	//以user_session作为会话名字，获取一个session对象
	sess, _ := session.Get(key, ctx)
	//设置会话参数
	sess.Options = &sessions.Options{
		Path:     "/",  //所有页面都可以访问会话数据
		MaxAge:   86400 * 30,   //会话有效期，单位秒
		HttpOnly: true,
	}
	//记录会话数据, sess.Values 是map类型，可以记录多个会话数据
	sess.Values[key] = value
	//保存用户会话数据
	sess.Save(ctx.Request(), ctx.Response())
}

//删除
func DelSession(ctx echo.Context, key string) {
	SetSession(ctx, key, nil)
}
