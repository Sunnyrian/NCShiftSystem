package routers

import (
	//"NCShiftSystem/web/controllers/admin"
	"NCShiftSystem/gin/controllers/admin"
	"NCShiftSystem/gin/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine)  {
	//路由分组 第二个参数可以设置该分组的所有中间件
	adminRouters := r.Group("/adminApi",middlewares.JwtAdminAuth())
	{
		//adminRouters.GET("/",admin.IndexController{}.Index)
		adminRouters.POST("/shift",admin.Controller{}.Shift)
		adminRouters.GET("/NAList",admin.Controller{}.GetNAList)
	}
}
