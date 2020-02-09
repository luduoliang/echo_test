package api

import (
	"echo_test/models"
	"echo_test/common"
	"github.com/labstack/echo"
)

//资讯分类列表，参数：
func CategoryList(c echo.Context) error  {
	list := models.GetAllCategoryList()
	return c.JSON(200, common.JsonData(true, list, "操作成功"))
}