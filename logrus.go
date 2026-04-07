package logsdk

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

// Loggers 定时日志logrus
type Loggers struct {
	logger *logrus.Logger
	module string
}

// newLogger 创建新的logger
func newLogger(logFormat string) *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetLevel(logrus.TraceLevel)

	var formatter logrus.Formatter
	timestampFormat := "2006-01-02 15:04:05.999999999"
	if logFormat == "text" {
		formatter = &logrus.TextFormatter{
			TimestampFormat: timestampFormat,
		}
	} else {
		formatter = &logrus.JSONFormatter{
			TimestampFormat:   timestampFormat,
			DisableHTMLEscape: true,
		}
	}
	logger.SetFormatter(formatter)

	// 设置日志打印位置
	logger.SetFormatter(
		&Caller{
			Formatter: formatter,
		},
	)
	return logger
}

// NewLogrus 日志初始化
func NewLogrus(opts Options) Logger {
	l := &Loggers{
		logger: newLogger(opts.LogFormat),
		module: opts.Module,
	}
	return l
}

// Debug logs a message at level Debug on the standard logger
func (l Loggers) Debug(ctx context.Context, message string, fields ...Fields) {
	l.logger.WithContext(ctx).WithFields(l.prepare(ctx, fields...)).Debug(message)
}

// Error logs a message at level Error on the standard logger
func (l Loggers) Error(ctx context.Context, message string, fields ...Fields) {
	l.logger.WithContext(ctx).WithFields(l.prepare(ctx, fields...)).Error(message)
}

// Info logs a message at level Info on the standard logger
func (l Loggers) Info(ctx context.Context, message string, fields ...Fields) {
	l.logger.WithContext(ctx).WithFields(l.prepare(ctx, fields...)).Info(message)
}

// Warn logs a message at level Warning on the standard logger
func (l Loggers) Warn(ctx context.Context, message string, fields ...Fields) {
	l.logger.WithContext(ctx).WithFields(l.prepare(ctx, fields...)).Debug(message)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1
func (l Loggers) Fatal(ctx context.Context, message string, fields ...Fields) {
	l.logger.WithContext(ctx).WithFields(l.prepare(ctx, fields...)).Fatal(message)
}

// Panic logs a message at level Panic on the standard logger
func (l Loggers) Panic(ctx context.Context, message string, fields ...Fields) {
	l.logger.WithContext(ctx).WithFields(l.prepare(ctx, fields...)).Panic(message)
}

// 日志打印前做数据准备和补充
func (l Loggers) prepare(ctx context.Context, fields ...Fields) logrus.Fields {
	// todo 彭亚川从ctx中获取上下文信息
	out := make(logrus.Fields, len(fields))
	out["module"] = l.module
	for _, v := range fields {
		vv := v.format()
		for k, v := range vv {
			out[k] = v
		}
	}
	return out
}
