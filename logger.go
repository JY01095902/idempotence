package idempotence

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
}

var _logger *zap.SugaredLogger
var loggerOnce sync.Once

func NewLogger() *zap.SugaredLogger {
	loggerOnce.Do(func() {
		config := zap.NewProductionConfig()
		config.Encoding = "console"
		config.EncoderConfig.CallerKey = ""
		config.EncoderConfig.StacktraceKey = ""
		config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		l, _ := config.Build()

		_logger = l.Sugar()
	})

	return _logger
}
