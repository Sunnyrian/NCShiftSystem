package routers

import (
	"NCShiftSystem/gin/controllers/portal"
	"github.com/gin-gonic/gin"
)

func PortalRouterInit(r *gin.Engine) {
	portalRouters := r.Group("/portal")
	{
		portalRouters.GET("/checkExist", portal.Controller{}.CheckUserExist)
	}
}