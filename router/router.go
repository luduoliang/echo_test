package router

import (
	"echo_test/middleware"
	"echo_test/controllers/admin"
	"echo_test/controllers/api"
	"github.com/labstack/echo"
	echo_middleware "github.com/labstack/echo/middleware"
)

//初始化路由
func InitRouter(e *echo.Echo){

	//登录后台
	e.Any("/admin/login", admin.Login)
	//验证码
	e.Any("/captcha", admin.Captcha)
	//退出登录
	e.Any("/admin/logout", admin.Logout)
	//后台路由，使用检查是否登录中间间
	admin_router := e.Group("/admin", middleware.CheckAdminLogin)
	{
		//后台首页
		admin_router.Any("/", admin.Index)
		//后台首页信息
		admin_router.Any("/info", admin.Info)
		//上传图片
		admin_router.Any("/file/upload", admin.UploadFile)

		//系统用户
		admin_router.Any("/admin/list", admin.AdminList)
		admin_router.Any("/admin/add", admin.AdminAdd)
		admin_router.Any("/admin/edit", admin.AdminEdit)
		admin_router.Any("/admin/delete", admin.AdminDelete)
		admin_router.Any("/admin/forbid", admin.AdminForbid)
		//资讯分类
		admin_router.Any("/category/list", admin.CategoryList)
		admin_router.Any("/category/add", admin.CategoryAdd)
		admin_router.Any("/category/edit", admin.CategoryEdit)
		admin_router.Any("/category/delete", admin.CategoryDelete)
		//资讯管理
		admin_router.Any("/news/list", admin.NewsList)
		admin_router.Any("/news/add", admin.NewsAdd)
		admin_router.Any("/news/edit", admin.NewsEdit)
		admin_router.Any("/news/delete", admin.NewsDelete)
	}


	//接口路由分组，使用设置跨域中间间
	api_router := e.Group("/api", echo_middleware.CORS())
	{
		//资讯分类列表，参数：
		api_router.Any("/category/list", api.CategoryList)
		//资讯列表，参数：page,per_page,category_id(可传)
		api_router.Any("/news/list", api.NewsList)
		//资讯详情，参数：参数：id
		api_router.Any("/news/detail", api.NewsDetail)



	}
}