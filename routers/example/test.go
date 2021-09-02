package example

import (
	"github.com/gin-gonic/gin"
	v1 "zhangyudevops.com/api/v1"
)

type TestRouter struct {
}

func (t *TestRouter) InitTestRouter(Router *gin.RouterGroup) {
	testRouter := Router.Group("test")
	var exaTestApi = v1.ApiGroupApp.TestApiGroup.BooksApi
	{
		testRouter.POST("register", exaTestApi.RegistryUser)
		testRouter.POST("login", exaTestApi.Login)
		testRouter.POST("deleteLoginUser", exaTestApi.DeleteLoginUser)
	}
}
