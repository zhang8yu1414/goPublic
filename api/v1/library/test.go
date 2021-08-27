package library

import (
	"github.com/gin-gonic/gin"
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
	userName := c.Request.FormValue("userName")
	password := c.Request.FormValue("password")
	phoneNumber := c.Request.FormValue("phoneNumber")

	password = utils.Md5Reverse(password)
	if err := LibService.CreateLoginUser(userName, password, phoneNumber); err != nil {
		response.FailWithMessage("创建用户失败", c)
		return
	}
	response.OKWithMessage("用户创建成功", c)
}