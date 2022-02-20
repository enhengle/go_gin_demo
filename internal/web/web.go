package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin/internal/exception"
)

/**
 * @Author:lingwang
 * @File :方法层
 * @Description:
 * @Version: 1.0.0
 * @Date :2021/12/27 11:36
 */

/*swagger注释*/
//Get请求
// GetUserInfo 获取用户信息接口
// @Tags base
// @Summary  获取用户信息接口
// @Produce json
// @Param   token     header    string     true     "登录token"
// @Success 200 {object} exception.Response{}
// @Router /WebOne [get]
func WebOne(c *gin.Context) {
	name := c.Query("name")
	fmt.Println(name)
	exception.SuccessResponse(c, name)
}

//Post请求：form
func WebTwo(c *gin.Context) {
	two := c.PostForm("two")
	exception.SuccessResponse(c, two)
}

//json请求结构
type WebThreeS struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
	EeE    string `json:"ee_e"`
}

//绑定json参数
func WebThree(c *gin.Context) {
	var webThreeS WebThreeS
	c.BindJSON(&webThreeS)
	exception.SuccessResponse(c, webThreeS)
}

//form请求结构
type WebThreeS2 struct {
	Title  string `form:"title"`
	Status int    `form:"status"`
	EeE    string `form:"ee_e"`
}

//绑定form参数
func WebFour(c *gin.Context) {
	var webThreeS WebThreeS2
	c.Bind(&webThreeS)
	exception.SuccessResponse(c, webThreeS)
}

//绑定json和form结构，且校验
type WebFourS struct {
	Title  string ` form:"title"	json:"title"	binding:"required"		 	validate:"required" `
	Status int    ` form:"status"	json:"status"	binding:"required,gte=0"	validate:"required" `
}

//绑定且校验
func WebFive(c *gin.Context) {
	var webFours WebFourS
	var webFours2 WebFourS
	err := c.ShouldBind(&webFours)
	err2 := c.ShouldBindJSON(&webFours2)
	if err == nil && err2 == nil {
		exception.SuccessResponse(c, webFours)
		exception.SuccessResponse(c, webFours2)
		return
	}
	exception.Error(c, exception.ResponseError("程序出错"))
	return
}
