package system

import (
	"github.com/gin-gonic/gin"
	v1 "zhangyudevops.com/api/v1"
)

type BaseRouter struct {

}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup)  {
	baseRouter := Router.Group("base")
	var exaUserLoginApi = v1.ApiGroupApp.TestApiGroup.UserApi
	{
		// login
		baseRouter.POST("register", exaUserLoginApi.RegistryUser)
		baseRouter.POST("login", exaUserLoginApi.Login)
		baseRouter.POST("deleteLoginUser", exaUserLoginApi.DeleteLoginUser)

	}
}