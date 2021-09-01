package library

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"zhangyudevops.com/core/books"
	"zhangyudevops.com/global"
	"zhangyudevops.com/model/common/response"
	"zhangyudevops.com/utils"
)

type BooksApi struct {
	
}

func (t *BooksApi) TestInfo(c *gin.Context)  {
	response.OKWithDetailed("Hello World!", "成功", c)
}

// CreateLoginUser 创建登录用户
func (t BooksApi) CreateLoginUser(c *gin.Context)  {
	data := books.GetLoginPostData(c)
	if err := LibService.CreateLoginUser(data); err != nil {
		global.GVA_LOG.Error("登录用户注册失败", zap.Any("err", err))
		response.FailWithMessage("创建用户失败", c)
		return
	}
	response.OKWithMessage("用户创建成功", c)
}

// GetLoginUser 查询登录用户
func (t BooksApi) GetLoginUser(c *gin.Context) {
	userName := c.Request.FormValue("userName")
	password := c.Request.FormValue("password")

	password = utils.Md5Reverse(password)

	result := LibService.SelectLoginUser(userName, password)
	if result.Error != nil {
		global.GVA_LOG.Error("数据查询失败", zap.Any("err", result.Error))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OKWithMessage("用户查询成功", c)
	}
}

// DeleteLoginUser 删除登录用户
func (t BooksApi) DeleteLoginUser(c *gin.Context)  {
	userName := c.Request.FormValue("userName")
	result := LibService.DeleteLoginUser(userName)
	if result.Error == nil {
		global.GVA_LOG.Info("登录用户删除成功", zap.String("username", userName))
		response.OKWithMessage("登录用户删除成功", c)
	} else {
		response.FailWithMessage("删除用户失败", c)
	}

}