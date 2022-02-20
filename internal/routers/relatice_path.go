package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go_gin/docs"
	"go_gin/internal/web"
	"net/http"
)

/**
 * @Author:lingwang
 * @File :初始化路由
 * @Description:
 * @Version: 1.0.0
 * @Date :2021/12/27 11:36
 */
func Load(r *gin.Engine) {
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	r.Any("/WebOne", web.WebOne)
	r.Any("/WebTwo", web.WebTwo)
	r.Any("/WebThree", web.WebThree)
	r.Any("/WebFour", web.WebFour)
	r.Any("/WebFive", web.WebFive)

	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
