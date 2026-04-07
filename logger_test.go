package logsdk

import (
	"context"
	"testing"
	"time"
)

func Test(t *testing.T) {
	Warn(context.TODO(), "测试代码", Any("1", "2"))

	time.Sleep(time.Second * 3)
}
