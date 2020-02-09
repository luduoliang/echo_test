package common

import (
	"io"
	"github.com/labstack/echo"
	"html/template"
)

//自定义的模版引擎struct
type Template struct {
	Templates *template.Template
}

//实现接口，Render函数
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	//调用模版引擎渲染模版
	return t.Templates.ExecuteTemplate(w, name, data)
}

//初始化模板
func InitTemplate(e *echo.Echo){
	//初始化模版引擎
	t := &Template{
		Templates: template.Must(template.ParseGlob("views/*/*/*")),
	}
	//向echo实例注册模版引擎
	e.Renderer = t
}