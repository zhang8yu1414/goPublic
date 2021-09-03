package library

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"zhangyudevops.com/global"
	"zhangyudevops.com/model/books"
	"zhangyudevops.com/model/books/request"
	"zhangyudevops.com/model/common/response"
)

type BooksApi struct {

}


// AddStudent 添加学生信息
func (t *BooksApi) AddStudent(c *gin.Context) {
	var s request.Students
	_ = c.ShouldBindJSON(&s)

	student := &books.Students{
		Name:         s.Name,
		Gender:       s.Gender,
		Age:          s.Age,
		Phone:        s.Phone,
	}

	if err := LibService.CreateStudent(*student); err != nil {
		global.GVA_LOG.Error("新增学生失败", zap.Any("err", err))
		response.FailWithMessage("新增学生失败", c)
		return
	}
	response.OKWithMessage("学生创建成功", c)
}

func (t *BooksApi) QueryStudentsInfo(c *gin.Context) {
	if err, result := LibService.GetAllStudentsInfo(); err != nil {
		global.GVA_LOG.Error("查询学生失败", zap.Any("err", err))
		response.FailWithMessage("查询学生失败", c)
		return
	} else {
		fmt.Println(result)
		response.OKWithDetailed(result, "学生信息查询成功", c)
	}
}
