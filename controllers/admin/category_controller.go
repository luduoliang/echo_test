package admin

import (
	"echo_test/models"
	"echo_test/library"
	"echo_test/common"
	"net/http"
	"strconv"
	"github.com/labstack/echo"
)

//列表
func CategoryList(c echo.Context) error {
	page_temp := c.FormValue("p")
	page, _ := strconv.Atoi(page_temp)
	list, nums := models.GetCategoryListPaginator(page)
	paginator := library.NewPaginator(c.Request(), models.PAGE_LIMIT, nums)

	return c.Render(http.StatusOK, "admin/category/list.tmpl",  map[string]interface{}{
		"list": list,
		"Page": paginator,
	})
}

//添加
func CategoryAdd(c echo.Context) error  {
	if(c.Request().Method == "POST"){
		name := c.FormValue("name")
		weight_temp := c.FormValue("weight")
		weight, _ := strconv.Atoi(weight_temp)
		info := models.Category{}
		info.Name = name
		info.Weight = weight
		info = models.CreateCategory(info)
		return c.JSON(200, common.JsonData(true, info, "操作成功"))
	}else{
		return c.Render(http.StatusOK, "admin/category/add.tmpl", nil)
	}

}

//编辑
func CategoryEdit(c echo.Context) error  {
	id_temp := c.FormValue("id")
	id, _ := strconv.Atoi(id_temp)
	info := models.GetCategory(id)
	if(c.Request().Method == "POST"){
		name := c.FormValue("name")
		weight_temp := c.FormValue("weight")
		weight, _ := strconv.Atoi(weight_temp)

		info.Name = name
		info.Weight = weight
		info = models.UpdateCategory(info)
		return c.JSON(200, common.JsonData(true, info, "操作成功"))
	}else{
		return c.Render(http.StatusOK, "admin/category/edit.tmpl", map[string]interface{}{
			"info": info,
		})
	}
}


//删除
func CategoryDelete(c echo.Context) error  {
	if(c.Request().Method == "POST"){
		id_temp := c.FormValue("id")
		id, _ := strconv.Atoi(id_temp)
		info := models.DeleteCategory(id)
		return c.JSON(200, common.JsonData(true, info, "操作成功"))
	}
	return nil
}

