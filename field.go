package logsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"reflect"
	"time"
)

// Fields fields
type Fields map[string]interface{}

// New 构建Fields对象
func New() Fields {
	return Fields{}
}

// WithErr 快捷设置错误日志
func (f Fields) WithErr(err error) Fields {
	f["error"] = err
	return f
}

// WithErrStack 快捷设置错误日志
func (f Fields) WithErrStack(errStack string) Fields {
	f["errorsStack"] = errStack
	return f
}

// WithAny 快捷设置任意日志
func (f Fields) WithAny(k string, v interface{}) Fields {
	f[k] = v
	return f
}

// format 格式化数据
func (f Fields) format() Fields {
	out := New()
	for k, v := range f {
		switch v.(type) {
		case error:
			out[k] = v

			// go-errors 追加错误栈
			if er, ok := v.(*errors.Error); ok {
				out.WithErrStack(er.ErrorStack())
				break // 跳出本次循环
			}

			// 其他errors类型
			e := v.(error)
			base := e.Error()
			e = errors.Wrap(e, 1)
			verbose := fmt.Sprintf("%+v", e)
			if verbose != base && k == "error" {
				out.WithAny("error", verbose)
			}
			if verbose != base && k != "error" {
				out.WithAny("errorsStack", verbose)
			}

		case []byte:
			out.WithAny(k, string(v.([]byte)))
		case time.Duration:
			out.WithAny(k, v.(time.Duration).String())
		case time.Time:
			out.WithAny(k, v.(time.Time).Format(time.RFC3339))
		case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128: // 基础类型以原样格式输出
			out.WithAny(k, v)
		default:
			if reflect.ValueOf(v).Kind() == reflect.String {
				out.WithAny(k, v)
			} else {
				fStr, _ := json.Marshal(v) // 其他类型统一转换为json字符串
				out.WithAny(k, string(fStr))
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
