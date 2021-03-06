package library

import (
	"github.com/gin-gonic/gin"
	v1 "zhangyudevops.com/api/v1"
)

type TestRouter struct {
}

func (t *TestRouter) InitStudentRouter(Router *gin.RouterGroup) {
	testRouter := Router.Group("test")
	var exaTestApi = v1.ApiGroupApp.TestApiGroup.BooksApi
	{
		//	student
		testRouter.POST("createStudent", exaTestApi.AddStudent)
		testRouter.GET("queryStudent", exaTestApi.QueryStudentsInfo)
	}
}
