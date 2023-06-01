package logdk

import (
	"context"
	"fmt"
	"runtime/debug"
)

// 默认logger
var defaultLogger = NewLogrus(Options{
	Module:       "demo",
	ToFieldsFunc: DefaultToFieldsFunc,
})

// SetLogger 设置logger
func SetLogger(logger Logger) {
	defaultLogger = logger
}

// Debug logs a message at level Debug on the standard logger
func Debug(ctx context.Context, message string, fields Fields) {
	defaultLogger.Debug(ctx, message, fields)
}

// Error logs a message at level Error on the standard logger
func Error(ctx context.Context, message string, fields Fields) {
	defaultLogger.Error(ctx, message, fields)
}

// Info logs a message at level Info on the standard logger
func Info(ctx context.Context, message string, fields Fields) {
	defaultLogger.Info(ctx, message, fields)
}

// Warning logs a message at level Warning on the standard logger
func Warning(ctx context.Context, message string, fields Fields) {
	defaultLogger.Warning(ctx, message, fields)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1
func Fatal(ctx context.Context, message string, fields Fields) {
	defaultLogger.Fatal(ctx, message, fields)
}

// Panic panic 日志，调用此方法后会触发panic捕获并且打印日志，外部无需再次捕获Panic 若未触发Panic不会打印日志
func Panic(ctx context.Context) {
	if err := recover(); err != nil {
		er := fmt.Errorf("[panic error] %+v", err)
		defaultLogger.Panic(ctx, "应用触发Panic", New().WithErr(er).WithAny("stack", debug.Stack()))
	}
}
