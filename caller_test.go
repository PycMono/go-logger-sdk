package logsdk

import (
	"testing"

	"github.com/sirupsen/logrus"
)

type noopFormatter struct{}

func (noopFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte("ok"), nil
}

func TestCallerFormatAddsCallerField(t *testing.T) {
	c := &Caller{Formatter: noopFormatter{}}
	entry := &logrus.Entry{Data: logrus.Fields{"k": "v"}}

	_, err := c.Format(entry)
	if err != nil {
		t.Fatalf("unexpected format error: %v", err)
	}

	if _, ok := entry.Data[CallerKey]; !ok {
		t.Fatalf("caller field missing after format")
	}
}

