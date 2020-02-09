package admin

import (
	"strconv"
	"time"
	"os"
	"echo_test/common"
	"net/http"
	"github.com/labstack/echo"
	"io"
	"encoding/json"
	"echo_test/models"
)

//首页
func Index(c echo.Context) error {
	admin_info_str := common.GetSession(c, "admin_info")
	var admin_info models.Admin
	json.Unmarshal([]byte(admin_info_str.(string)), &admin_info)
	return c.Render(http.StatusOK, "admin/index/index.tmpl", map[string]interface{}{
		"admin_info": admin_info,
	})
}

//首页信息
func Info(c echo.Context) error  {
	return c.Render(http.StatusOK, "admin/index/info.tmpl", nil)
}

//上传文件
func UploadFile(c echo.Context) error  {

	file, form_err := c.FormFile("pic")
	if form_err != nil {
		return c.JSON(200, common.JsonData(false, "", form_err.Error()))
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(200, common.JsonData(false, "", err.Error()))
	}

	fname := file.Filename

	year := strconv.Itoa(time.Now().Year())
	month := time.Now().Month().String()
	day := strconv.Itoa(time.Now().Day())
	//创建目录
	dirpath := "static/upload/"+year+"/"+month+"/"+day
	dir_err := os.MkdirAll(dirpath, os.ModePerm)
	if dir_err != nil {
		return c.JSON(200, common.JsonData(false, "", dir_err.Error()))
	}

	//上传文件相对路劲
	file_path := dirpath + "/" + fname

	out, open_err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0666)
	defer out.Close()

	if open_err != nil {
		return c.JSON(200, common.JsonData(false, "", open_err.Error()))
	}

	// Copy
	if _, err = io.Copy(out, src); err != nil {
		return c.JSON(200, common.JsonData(false, "", err.Error()))
	}

	scheme := "http://"
	if c.Request().TLS != nil {
		scheme = "https://"
	}

	all_file_path := scheme + c.Request().Host + "/" + file_path
	upload_info := make(map[string]string)
	upload_info["pic_url"] = all_file_path

	return c.JSON(200, common.JsonData(true, upload_info, "上传成功"))
}


