package initial

import (
	"github.com/gin-gonic/gin"
	"zhangyudevops.com/global"
	"zhangyudevops.com/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	exampleRouter := router.RouterGroupApp.Example
	PublicGroup := Router.Group("")
	{
		exampleRouter.InitTestRouter(PublicGroup)
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
