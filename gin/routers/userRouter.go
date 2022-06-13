package routers

import (
	"NCShiftSystem/gin/controllers/user"
	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	userRouters := r.Group("/userApi")
	{
		userRouters.GET("/getUserInfo", user.Controller{}.GetUserInformation)
	}
}