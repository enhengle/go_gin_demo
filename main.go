package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin/internal/config"
	"go_gin/internal/exception"
	"go_gin/internal/routers"
	"go_gin/internal/xlogger"
	"go_gin/support/genv"
)

var configFile = flag.String("develop", ".env.develop", "the config file")

/**
 * @Author:lingwang
 * @File :主方法
 * @Description:
 * @Version: 1.0.0
 * @Date :2021/12/27 11:36
 */
func main() {
	//读取配置文件
	flag.Parse()
	var c config.Config
	err := genv.EnvLoader(*configFile, &c)
	if err != nil {
		xlogger.GatewayLogger.Error(err.Error())
	}

	//初始化路由
	r := gin.Default()
	routers.Load(r)

	//全局异常处理
	r.NoMethod(exception.HandlerNotFound)
	r.NoRoute(exception.HandlerNotFound)
	r.Use(exception.ErrorHandler())

	//设置端口
	post := fmt.Sprintf(":%d", c.Port)

	r.Run(post)

}
