package exception

import "github.com/gin-gonic/gin"

/**
 * @Author:lingwang
 * @File :全局异常捕捉
 * @Description:
 * @Version: 1.0.0
 * @Date :2021/12/27 11:36
 */
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if Err := recover(); Err != nil {
			var err *Response
			if e, ok := Err.(*Response); ok {
				err = e
			} else if e, ok := Err.(error); ok {
				err = ResponseError(e.Error())
			} else {
				err = ServerError
			}
			c.JSON(err.Code, err)
			return
		}
		c.Next()
	}
}
