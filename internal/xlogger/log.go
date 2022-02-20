package xlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/**
 * @Author:lingwang
 * @File :日志
 * @Description:
 * @Version: 1.0.0
 * @Date :2021/12/27 11:36
 */

var MainLogger *zap.Logger
var GatewayLogger *zap.Logger

func init() {
	MainLogger = NewLogger("./logs/main.log", zapcore.InfoLevel, 128, 30, 7, true, "Main")
	GatewayLogger = NewLogger("./logs/gateway.log", zapcore.DebugLevel, 128, 30, 7, true, "Gateway")
}
