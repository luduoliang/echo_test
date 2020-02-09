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
func AdminList(c echo.Context) error {
	page_temp := c.FormValue("p")
	page, _ := strconv.Atoi(page_temp)
	list, nums := models.GetAdminListPaginator(page)
	paginator := library.NewPaginator(c.Request(), models.PAGE_LIMIT, nums)

	return c.Render(http.StatusOK, "admin/admin/list.tmpl", map[string]interface{}{
		"list": list,
		"Page": paginator,
	})
}

//添加
func AdminAdd(c echo.Context) error  {
	if(c.Request().Method == "POST"){
		username := c.FormValue("username")
		password := c.FormValue("password")

		info := models.Admin{}
		info.Username = username
		info.Password = password
		info = models.CreateAdmin(info)

		return c.JSON(200, common.JsonData(true, info, "操作成功"))
	}else{
		return c.Render(http.StatusOK, "admin/admin/add.tmpl", nil)
	}

}

//编辑
func AdminEdit(c echo.Context) error  {
	id_temp := c.FormValue("id")
	id, _ := strconv.Atoi(id_temp)
	info := models.GetAdmin(id)
	if(c.Request().Method == "POST"){
		username := c.FormValue("username")
		password := c.FormValue("password")

		info.Username = username
		info.Password = password
		info = models.UpdateAdmin(info)
		return c.JSON(200, common.JsonData(true, info, "操作成功"))
	}else{
		return c.Render(http.StatusOK, "admin/admin/edit.tmpl", map[string]interface{}{
			"info": info,
		})
	}
}


//删除
func AdminDelete(c echo.Context) error  {
	if(c.Request().Method == "POST"){
		id_temp := c.FormValue("id")
		id, _ := strconv.Atoi(id_temp)

		info := models.DeleteAdmin(id)
		return c.JSON(200, common.JsonData(true, info, "操作成功"))
	}

	return nil
}

//禁用/恢复
func AdminForbid(c echo.Context) error  {
	if(c.Request().Method == "POST"){
		id_temp := c.FormValue("id")
		id, _ := strconv.Atoi(id_temp)
		info := models.GetAdmin(id)

		if(info.Status == 1){
			info.Status = 2
		}else if(info.Status == 2){
			info.Status = 1
		}

		info = models.UpdateAdminInfo(info)
		return c.JSON(200, common.JsonData(true, info, "操作成功"))
	}
	return nil
}

