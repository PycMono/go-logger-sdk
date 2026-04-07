# go-logger-sdk

基于 `logrus` 的轻量日志封装，提供统一字段格式化、模块标识、调用位置（caller）和上下文字段注入能力。

## 安装

```bash
go get github.com/PycMono/go-logger-sdk
```

## 快速开始

```go
package main

import (
	"context"
	"errors"

	logsdk "github.com/PycMono/go-logger-sdk"
)

func main() {
	logger := logsdk.NewLogrus(logsdk.Options{
		LogFormat: "json", // "json" 或 "text"
		Module:    "order-service",
	})

	ctx := context.Background()
	logger.Info(ctx, "service started", logsdk.Any("port", 8080))
	logger.Error(ctx, "create order failed", logsdk.Err(errors.New("db timeout")))
}
```

如果你希望使用包级函数（`logsdk.Info`、`logsdk.Error` 等），可以先设置默认 logger：

```go
logsdk.SetLogger(logsdk.NewLogrus(logsdk.Options{
	LogFormat: "json",
	Module:    "payment-service",
}))

logsdk.Info(context.Background(), "ready", logsdk.Any("version", "v1.0.0"))
```

## 字段 API

`Fields` 是 `map[string]interface{}` 的别名，支持多个字段对象一起传入，后传入同名 key 会覆盖前一个。

- `logsdk.N()`：创建空字段对象
- `logsdk.Any(key, value)`：快速创建单字段
- `logsdk.Err(err)`：写入 `error` 字段
- `logsdk.ErrStack(stack)`：写入 `errorsStack` 字段

示例：

```go
logsdk.Info(ctx, "query ok",
	logsdk.Any("user_id", 1001),
	logsdk.Any("cost", 12.3),
	logsdk.Any("at", time.Now()),
)
```

字段会在内部自动标准化：

- `[]byte` -> 字符串
- `time.Duration` -> `String()`
- `time.Time` -> RFC3339
- 基础数值/布尔类型 -> 原样输出
- 复杂对象 -> JSON 字符串
- `error` -> 保留 `error` 字段；若是 `go-errors`，会追加 `errorsStack`

## 配置项

`Options`:

- `LogFormat string`：日志格式，`"text"` 或 `"json"`（默认走 JSON 分支）
- `Module string`：固定注入到每条日志的 `module` 字段
- `ToFieldsFunc ToFieldsFunc`：从 `context.Context` 注入/加工字段

`ToFieldsFunc` 例子：

```go
type traceKey struct{}

logger := logsdk.NewLogrus(logsdk.Options{
	Module: "gateway",
	ToFieldsFunc: func(ctx context.Context, fields logsdk.Fields) logsdk.Fields {
		if traceID, ok := ctx.Value(traceKey{}).(string); ok && traceID != "" {
			fields["trace_id"] = traceID
		}
		return fields
	},
})
```

## Panic 用法

### 1) 主动记录 panic 级别日志

```go
logger.Panic(ctx, "unexpected state", logsdk.Any("order_id", 123))
```

### 2) 在 `defer` 中捕获并打印 panic（包级函数）

```go
func run(ctx context.Context) {
	defer logsdk.Panic(ctx)

	// your code...
}
```

`logsdk.Panic(ctx)` 会在 `recover()` 到异常时输出 panic 日志，并携带堆栈字段。

## 说明

- `Warn` 会按 warning 级别输出。
- 每条日志默认包含 `caller` 与 `module` 字段。
