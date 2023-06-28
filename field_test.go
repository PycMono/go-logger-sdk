package logsdk

import (
	"context"
	"github.com/go-errors/errors"
	"testing"
	"time"
)

type PreFields struct {
}

func TestEncodeFields(t *testing.T) {
	fields := New()
	var preNil PreFields
	preStruct := PreFields{}
	var fv float32 = 1.1
	fields.WithAny("1", "2").
		WithAny("1", 3).
		WithAny("int8", int8(1)).
		WithAny("int64", int64(5)).
		WithAny("int32", int32(6)).
		WithAny("nil", nil).
		WithErr(errors.Errorf("失败")).
		WithAny("time", time.Now()).
		WithErrStack(errors.New("stack").ErrorStack()).
		WithAny("preNil", preNil).
		WithAny("preStruct", preStruct).
		WithAny("byte", []byte("你好")).
		WithAny("float32", fv).
		WithAny("map", map[string]any{"abc": 1}).
		WithAny("uint", uint16(1)).
		WithAny("struct", struct{}{}).
		WithAny("array", []int{1, 2, 3, 4})
	Info(context.TODO(), "test", fields)

	//fields = New()
	//fields.WithAny("func", func() int {
	//	return 1
	//}).WithAny("err", fmt.Errorf("err%+v", "errorsss")).WithErr(errors.New("失败"))
	//Warning(context.TODO(), "test2", fields)

	//fields = New()
	//fields.WithAny("err", fmt.Errorf("err%+v", "errorsss"))
	//Warning(context.TODO(), "test2", fields)
}
