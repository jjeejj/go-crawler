package test

import (
	"testing"

	"github.com/jjeejj/go-crawler/log"
	"go.uber.org/zap/zapcore"
)

func TestLogFilePlugin(t *testing.T) {
	filePlugin, c := log.NewFilePlugin("./log.txt", zapcore.InfoLevel)
	defer c.Close()
	// stdoutPlugin := log.NewStdoutPlugin(zapcore.InfoLevel)
	logger := log.NewLogger(filePlugin)
	logger.Info("log init end")
}
