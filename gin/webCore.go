package web

import (
	"NCShiftSystem/gin/middlewares"
	"NCShiftSystem/gin/routers"
	// "NCShiftSystem/web/models"
	// "NCShiftSystem/web/routers"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// StartWeb 启动 gin
func StartWeb() {
	//创建一个默认的路由引擎
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./web/dist", false)))
	//r.Use(middlewares.Cors())
	routers.PortalRouterInit(r)
	r.Use(middlewares.JwtAuth())
	routers.UserRouterInit(r)
	routers.AdminRoutersInit(r)
	err := r.Run(":3500")
	if err != nil {
		fmt.Println(err)
	}
}


