package logsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-errors/errors"
)

// Fields fields
type Fields map[string]interface{}

// N 构建Fields对象
func N() Fields {
	return Fields{}
}

// Err 快捷设置错误日志
func Err(err error) Fields {
	return Fields{"error": err}
}

// ErrStack 快捷设置错误日志
func ErrStack(errStack string) Fields {
	return Fields{"errorsStack": errStack}
}

func Any(k string, v interface{}) Fields {
	return Fields{k: v}
}

// format 格式化数据
func (f Fields) format() Fields {
	out := N()
	for k, v := range f {
		if v == nil {
			out[k] = nil
			continue
		}

		switch v.(type) {
		case error:
			out[k] = v

			// go-errors 追加错误栈
			if er, ok := v.(*errors.Error); ok {
				out["errorsStack"] = er.ErrorStack()
				break // 跳出本次循环
			}

			// 其他errors类型
			e := v.(error)
			base := e.Error()
			e = errors.Wrap(e, 1)
			verbose := fmt.Sprintf("%+v", e)
			if verbose != base && k == "error" {
				out["error"] = verbose
			}
			if verbose != base && k != "error" {
				out["errorsStack"] = verbose
			}

		case []byte:
			out[k] = string(v.([]byte))
		case time.Duration:
			out[k] = v.(time.Duration).String()
		case time.Time:
			out[k] = v.(time.Time).Format(time.RFC3339)
		case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128: // 基础类型以原样格式输出
			out[k] = v
		default:
			if s, ok := v.(string); ok {
				out[k] = s
			} else {
				fStr, _ := json.Marshal(v) // 其他类型统一转换为json字符串
				out[k] = string(fStr)
			}
		}
	}
	return out
}

// DefaultToFieldsFunc 默认注入器
func DefaultToFieldsFunc(ctx context.Context, fields Fields) Fields {
	// todo 补充ctx 获取上线文信息操作
	return fields
}

// ToFieldsFunc 从ctx中获取字段并且整合fields字段
type ToFieldsFunc func(ctx context.Context, fields Fields) Fields
