package routers

import (
	"NCShiftSystem/gin/controllers/portal"
	"github.com/gin-gonic/gin"
)

func PortalRouterInit(r *gin.Engine) {
	portalRouters := r.Group("/portalApi")
	{
		portalRouters.GET("/checkExist", portal.Controller{}.CheckUserExist)
		portalRouters.POST("/register", portal.Controller{}.Register)
		portalRouters.POST("/login", portal.Controller{}.Login)
		portalRouters.GET("/checkLogin", portal.Controller{}.CheckLogin)
	}
}