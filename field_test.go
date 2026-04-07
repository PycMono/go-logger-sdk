package logsdk

//
//import (
//	"context"
//	"github.com/go-errors/errors"
//	"testing"
//	"time"
//)
//
//type PreFields struct {
//}
//
//func TestEncodeFields(t *testing.T) {
//	fields := N()
//	var preNil PreFields
//	preStruct := PreFields{}
//	var fv float32 = 1.1
//	Any("1", "2").
//		Any("1", 3).
//		Any("int8", int8(1)).
//		Any("int64", int64(5)).
//		Any("int32", int32(6)).
//		Any("nil", nil).
//		Err(errors.Errorf("失败")).
//		Any("time", time.Now()).
//		ErrStack(errors.New("stack").ErrorStack()).
//		Any("preNil", preNil).
//		Any("preStruct", preStruct).
//		Any("byte", []byte("你好")).
//		Any("float32", fv).
//		Any("map", map[string]any{"abc": 1}).
//		Any("uint", uint16(1)).
//		Any("struct", struct{}{}).
//		Any("array", []int{1, 2, 3, 4})
//	Info(context.TODO(), "test", fields)
//
//	//fields = New()
//	//fields.Any("func", func() int {
//	//	return 1
//	//}).Any("err", fmt.Errorf("err%+v", "errorsss")).WithErr(errors.New("失败"))
//	//Warning(context.TODO(), "test2", fields)
//
//	//fields = New()
//	//fields.Any("err", fmt.Errorf("err%+v", "errorsss"))
//	//Warning(context.TODO(), "test2", fields)
//}
