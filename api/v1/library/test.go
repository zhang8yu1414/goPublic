package library

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"
	"zhangyudevops.com/core/books"
	"zhangyudevops.com/global"
	"zhangyudevops.com/middleware"
	modelBook "zhangyudevops.com/model/books"
	"zhangyudevops.com/model/common/response"
	sysReq "zhangyudevops.com/model/system/request"
	modelResponse "zhangyudevops.com/model/system/response"
	"zhangyudevops.com/utils"
)

type BooksApi struct {
}

func (t *BooksApi) TestInfo(c *gin.Context) {
	response.OKWithDetailed("Hello World!", "成功", c)
}

// RegistryUser 创建登录用户
func (t *BooksApi) RegistryUser(c *gin.Context) {
	// 单个取值，绑定结构体写法
	data := books.GetLoginPostData(c)

	// @todo: ShouldBindJSON 写法， 不成功， 需要继续研究
	//var loginUser sysReq.Register
	//_ = c.ShouldBindJSON(&loginUser)
	//user := &modelBook.BookLoginUser{
	//	UserName: loginUser.Username,
	//	Password: loginUser.Password,
	//	PhoneNumber: loginUser.Password,
	//}

	if err := LibService.CreateLoginUser(data); err != nil {
		global.GVA_LOG.Error("登录用户注册失败", zap.Any("err", err))
		response.FailWithMessage("创建用户失败", c)
		return
	}
	response.OKWithMessage("用户创建成功", c)
}

// Login 用户登录
func (t *BooksApi) Login(c *gin.Context) {
	userName := c.Request.FormValue("userName")
	password := c.Request.FormValue("password")
	password = utils.Md5Reverse(password)

	user, err := LibService.SelectLoginUser(userName, password)
	if err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		t.tokenNext(c, *user)
	}
}

func (t *BooksApi) tokenNext(c *gin.Context, user modelBook.BookLoginUser) {
	j := &middleware.JWT{
		SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey),
	}
	claims := sysReq.CustomClaims{
		Username:   user.UserName,
		BufferTime: global.GVA_CONFIG.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime,
			Issuer:    "TokyoHot",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OKWithDetailed(modelResponse.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// DeleteLoginUser 删除登录用户
func (t *BooksApi) DeleteLoginUser(c *gin.Context) {
	userName := c.Request.FormValue("userName")
	_, err := LibService.DeleteLoginUser(userName)
	if err == nil {
		global.GVA_LOG.Info("登录用户删除成功", zap.String("username", userName))
		response.OKWithMessage("登录用户删除成功", c)
	} else {
		response.FailWithMessage("删除用户失败", c)
	}

}
