package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Plugin = zapcore.Core

func NewLogger(plugin zapcore.Core, option ...zap.Option) *zap.Logger {
	return zap.New(plugin, append(defaultOption(), option...)...)
}

// NewStdoutPlugin 输出到控制台
func NewStdoutPlugin(enabler zapcore.LevelEnabler) Plugin {
	return newPlugin(zapcore.Lock(zapcore.AddSync(os.Stdout)), enabler)
}

// NewStderrPlugin 错误输出到控制台
func NewStderrPlugin(enabler zapcore.LevelEnabler) Plugin {
	return newPlugin(zapcore.Lock(zapcore.AddSync(os.Stderr)), enabler)
}

// NewFilePlugin 打印到文件中
func NewFilePlugin(filepath string, enabler zapcore.LevelEnabler) (Plugin, io.Closer) {
	writer := defaultLumberjackLogger()
	writer.Filename = filepath
	return newPlugin(zapcore.AddSync(writer), enabler), writer
}

// newPlugin 创建底层日志的 core
func newPlugin(write zapcore.WriteSyncer, enabler zapcore.LevelEnabler) Plugin {
	return zapcore.NewCore(defaultEncoder(), write, enabler)
}
