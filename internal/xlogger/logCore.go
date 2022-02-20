package xlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

/**
 * @Author:lingwang
 * @File :日志具体实现方法
 * @Description:
 * @Version: 1.0.0
 * @Date :2021/12/27 11:36
 */

/**
 * 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 */
func NewLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int,
	maxAge int, compress bool, serviceName string) *zap.Logger {
	core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)
	return zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", serviceName)))
}

/**
 * zapcore构造
 */
func newCore(filePath string, level zapcore.Level, maxSize int, maxBackup int,
	maxAge int, compress bool) zapcore.Core {
	//日志文件路径
	hook := lumberjack.Logger{
		Filename:   filePath,  //日志文件路径
		MaxSize:    maxAge,    //每个文件最大保存尺寸，单位：M
		MaxBackups: maxBackup, //日志最多保存备份
		MaxAge:     maxAge,    //文件最多保存多少天
		Compress:   compress,  //是否压缩
	}
	//设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "linenum",
		StacktraceKey:  "stacktrace",
		MessageKey:     "msg", //打印信息
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, //小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    //UTF 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, //全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           //编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), //打印到控制台和文件
		atomicLevel, //日志级别
	)
}
