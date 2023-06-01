package logdk

import (
	"context"
	"testing"
	"time"
)

func Test(t *testing.T) {
	Warning(context.TODO(), "测试代码", New().WithAny("1", "2"))

	time.Sleep(time.Second * 3)
}
