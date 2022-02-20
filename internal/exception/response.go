package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin/internal/xlogger"
	"net/http"
)

/**
 * @Author:lingwang
 * @File :返回结构
 * @Description:
 * @Version: 1.0.0
 * @Date :2021/12/27 11:36
 */
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	ServerError   = ResponseError("系统异常")
	NotFoundError = HandlerNotFound
)

func ResponseError(err string) *Response {
	return &Response{
		Code: 5050,
		Msg:  err,
		Data: nil,
	}
}

func HandlerNotFound(c *gin.Context) {
	err := Response{
		Code: c.Writer.Status(),
		Msg:  c.Request.RequestURI,
		Data: nil,
	}
	xlogger.MainLogger.Error(fmt.Sprintf("%d  %s", err.Code, err.Msg))
	c.JSON(err.Code, err)
	return
}

func SuccessResponse(c *gin.Context, data interface{}) {
	response := Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
	c.JSON(http.StatusOK, response)
}

func Error(c *gin.Context, response *Response) {
	c.JSON(http.StatusOK, &response)
}
