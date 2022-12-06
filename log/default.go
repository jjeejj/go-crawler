package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// DefaultOption 默认的日志配置项
func defaultOption() []zap.Option {
	// 判断，只有设置的当前等级 高于 DPanicLevel 才需要打印 堆栈信息
	var stacktraceLevel zap.LevelEnablerFunc = func(l zapcore.Level) bool {
		return l >= zapcore.DPanicLevel
	}
	return []zap.Option{
		zap.AddCallerSkip(1),
		zap.AddStacktrace(stacktraceLevel),
		zap.AddCaller(),
	}
}

// 1.不会自动清理backup
// 2.每200MB压缩一次，不按时间切割
func defaultLumberjackLogger() *lumberjack.Logger {
	return &lumberjack.Logger{
		MaxSize:   200,
		LocalTime: true,
		Compress:  true,
	}
}

// defaultEncoder 默认使用json
func defaultEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(defaultEncoderConfig())
}

func defaultEncoderConfig() zapcore.EncoderConfig {
	config := zap.NewProductionEncoderConfig()
	// config.EncodeLevel
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	return config
}
