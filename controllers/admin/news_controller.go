package admin

import (
	"echo_test/models"
	"echo_test/library"
	"echo_test/common"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

//列表
func NewsList(c echo.Context) error {
	page_temp := c.FormValue("p")
	page, _ := strconv.Atoi(page_temp)
	list, nums := models.GetNewsListPaginator(page)
	paginator := library.NewPaginator(c.Request(), models.PAGE_LIMIT, nums)

	return c.Render(http.StatusOK, "admin/news/list.tmpl", map[string]interface{}{
		"list": list,
		"Page": paginator,
	})
}

//添加
func NewsAdd(c echo.Context) error  {
	if(c.Request().Method == "POST"){
		category_id_temp :=  c.FormValue("category_id")
		category_id, _ := strconv.Atoi(category_id_temp)
		title :=  c.FormValue("title")
		description :=  c.FormValue("description")
		pic :=  c.FormValue("pic")
		content :=  c.FormValue("content")
		weight_temp := c.FormValue("weight")
		weight, _ := strconv.Atoi(weight_temp)

		info := models.News{}
		info.CategoryId = category_id
		info.Title = title
		info.Description = description
		info.Pic = pic
		info.Content = content
		info.Weight = weight

		info = models.CreateNews(info)
		return c.JSON(200, common.JsonData(true, info, "操作成功"))
	}else{
		category := models.GetAllCategoryList()

		return c.Render(http.StatusOK, "admin/news/add.tmpl", map[string]interface{}{
			"category": category,
		})
	}

}

//编辑
func NewsEdit(c echo.Context) error  {
	id_temp := c.FormValue("id")
	id, _ := strconv.Atoi(id_temp)
	info := models.GetNews(id)
	if(c.Request().Method == "POST"){
		category_id_temp :=  c.FormValue("category_id")
		category_id, _ := strconv.Atoi(category_id_temp)
		title :=  c.FormValue("title")
		description :=  c.FormValue("description")
		pic :=  c.FormValue("pic")
		content :=  c.FormValue("content")
		weight_temp := c.FormValue("weight")
		weight, _ := strconv.Atoi(weight_temp)

		info.CategoryId = category_id
		info.Title = title
		info.Description = description
		info.Pic = pic
		info.Content = content
		info.Weight = weight

		info = models.UpdateNews(info)
		return c.JSON(200, common.JsonData(true, info, "操作成功"))
	}else{
		category := models.GetAllCategoryList()

		return c.Render(http.StatusOK, "admin/news/edit.tmpl", map[string]interface{}{
			"category": category,
			"info": info,
		})
	}
}


//删除
func NewsDelete(c echo.Context) error  {
	if(c.Request().Method == "POST"){
		id_temp := c.FormValue("id")
		id, _ := strconv.Atoi(id_temp)

		info := models.DeleteNews(id)
		return c.JSON(200, common.JsonData(true, info, "操作成功"))
	}
	return nil
}
