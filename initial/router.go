package initial

import (
	"github.com/gin-gonic/gin"
	"zhangyudevops.com/global"
	"zhangyudevops.com/middleware"
	"zhangyudevops.com/routers"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	exampleRouter := routers.RouterGroupApp.Example
	BaseRouter := routers.RouterGroupApp.System
	PublicGroup := Router.Group("")
	{
		BaseRouter.InitBaseRouter(PublicGroup)
	}

	PrivateRouter := Router.Group("")
	PrivateRouter.Use(middleware.JWTAuth())
	{
		exampleRouter.InitStudentRouter(PrivateRouter)
	}

	global.GVA_LOG.Info("routers register success")
	return Router
}
