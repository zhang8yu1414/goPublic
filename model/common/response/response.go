package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 4
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func OK(c *gin.Context)  {
	Result(SUCCESS, map[string]interface{}{},"操作成功", c)
}

func OKWithMessage(message string,c *gin.Context)  {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OKWithData(data interface{}, c *gin.Context)  {
	Result(SUCCESS, data, "操作成功", c)
}

func OKWithDetailed(data interface{}, message string, c *gin.Context)  {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context)  {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context)  {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetail(data interface{}, message string, c *gin.Context)  {
	Result(ERROR, data, message, c)
}