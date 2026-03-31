package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// @logger-config
// 功能: 初始化zap日志
// 输出: stdout/stderr (kubectl logs可见)
// 格式: JSON (云原生标准)
// 时间: ISO8601格式
// 备注：符合12 Factor App规范，将日志输出到stdout/stderr，方便容器化部署
var log *zap.Logger

// @logger-init
func InitLogger() {
	// 配置生产环境日志
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}
	config.Encoding = "json"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 构建并设置全局日志
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	log = logger
	zap.ReplaceGlobals(logger)
}

// @logger-method
func Info(msg string, args ...zap.Field) {
	log.Info(msg, args...)
}

// @logger-method
func Warn(msg string, args ...zap.Field) {
	log.Warn(msg, args...)
}

// @logger-method
func Error(msg string, args ...zap.Field) {
	log.Error(msg, args...)
}

// @logger-method
func Fatal(msg string, args ...zap.Field) {
	log.Fatal(msg, args...)
}
