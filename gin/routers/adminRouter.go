package routers

import (
	//"NCShiftSystem/web/controllers/admin"
	"NCShiftSystem/gin/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine)  {
	//路由分组 第二个参数可以设置该分组的所有中间件
	adminRouters := r.Group("/adminApi")
	{
		//adminRouters.GET("/",admin.IndexController{}.Index)
		adminRouters.POST("/shift",admin.IndexController{}.Shift)
	}
}
