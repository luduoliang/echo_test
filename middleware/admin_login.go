package middleware

import (
	"github.com/labstack/echo"
	"echo_test/common"
	"net/http"
	"errors"
)

//中间件函数
func CheckAdminLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		admin_info := common.GetSession(c, "admin_info")
		if admin_info == nil {
			c.Redirect(http.StatusFound, "/admin/login")
			return errors.New("请登录")
		}
		//执行下一个中间件或者执行控制器函数, 然后返回执行结果
		return next(c)
	}
}
