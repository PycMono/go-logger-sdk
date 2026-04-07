package logsdk

import "context"

// Logger 日志接口
type Logger interface {
	// Debug logs a message at level Debug on the standard logger
	Debug(ctx context.Context, message string, fields ...Fields)
	// Error logs a message at level Error on the standard logger
	Error(ctx context.Context, message string, fields ...Fields)
	// Info logs a message at level Info on the standard logger
	Info(ctx context.Context, message string, fields ...Fields)
	// Warn logs a message at level Warning on the standard logger
	Warn(ctx context.Context, message string, fields ...Fields)
	// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1
	Fatal(ctx context.Context, message string, fields ...Fields)
	// Panic logs a message at level Panic on the standard logger
	Panic(ctx context.Context, message string, fields ...Fields)
}
