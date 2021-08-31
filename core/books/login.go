package books

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"zhangyudevops.com/global"
	"zhangyudevops.com/model/books"
	"zhangyudevops.com/utils"
)

func GetLoginPostData(c *gin.Context) (user books.BookLoginUser) {
	userName := c.Request.FormValue("userName")
	password := c.Request.FormValue("password")
	phoneNumber := c.Request.FormValue("phoneNumber")

	password = utils.Md5Reverse(password)
	user = books.BookLoginUser{
		UserName: userName,
		Password: password,
		PhoneNumber: phoneNumber,
	}
	global.GVA_LOG.Info("提交的user信息为", zap.Any("info", user))

	return user
}
