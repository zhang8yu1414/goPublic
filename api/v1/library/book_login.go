package library

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"
	"zhangyudevops.com/global"
	"zhangyudevops.com/middleware"
	modelBook "zhangyudevops.com/model/books"
	sysReq "zhangyudevops.com/model/books/request"
	modelResponse "zhangyudevops.com/model/books/response"
	"zhangyudevops.com/model/common/response"
	"zhangyudevops.com/model/system/request"
	"zhangyudevops.com/utils"
)

type UserApi struct {
}

// RegistryUser 创建登录用户
func (t *UserApi) RegistryUser(c *gin.Context) {

	// @todo: ShouldBindJSON 写法， 成功
	var loginUser sysReq.Register
	_ = c.ShouldBindJSON(&loginUser)
	user := &modelBook.BookLoginUser{
		UserName: loginUser.Username,
		Password: loginUser.Password,
		PhoneNumber: loginUser.PhoneNumber,
	}

	if err := LibService.CreateLoginUser(*user); err != nil {
		global.GVA_LOG.Error("登录用户注册失败", zap.Any("err", err))
		response.FailWithMessage("创建用户失败", c)
		return
	}
	response.OKWithMessage("用户创建成功", c)
}

// Login 用户登录
func (t *UserApi) Login(c *gin.Context) {
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

func (t *UserApi) tokenNext(c *gin.Context, user modelBook.BookLoginUser) {
	j := &middleware.JWT{
		SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey),
	}
	claims := request.CustomClaims{
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
func (t *UserApi) DeleteLoginUser(c *gin.Context) {
	userName := c.Request.FormValue("userName")
	_, err := LibService.DeleteLoginUser(userName)
	if err == nil {
		global.GVA_LOG.Info("登录用户删除成功", zap.String("username", userName))
		response.OKWithMessage("登录用户删除成功", c)
	} else {
		response.FailWithMessage("删除用户失败", c)
	}

}
