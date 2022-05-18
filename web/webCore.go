package web

import (
	"NCShiftSystem/web/models"
	"NCShiftSystem/web/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

// StartWeb 启动 gin
func StartWeb() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数 要放在加载模板之前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	r.LoadHTMLGlob("./web/templates/**/*")

	routers.AdminRoutersInit(r)
	err := r.Run(":80")
	if err != nil {
		fmt.Println(err)
	}
}