package logsdk

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestWarnUsesWarningLevel(t *testing.T) {
	l := NewLogrus(Options{Module: "test"}).(*Loggers)

	buf := bytes.NewBuffer(nil)
	l.logger.Out = buf

	l.Warn(context.Background(), "warn message", Any("k", "v"))

	logLine := buf.String()
	if !strings.Contains(logLine, `"level":"warning"`) {
		t.Fatalf("warn should emit warning level, got log: %s", logLine)
	}
}

func TestPrepareAppliesToFieldsFunc(t *testing.T) {
	l := NewLogrus(Options{
		Module: "test",
		ToFieldsFunc: func(ctx context.Context, fields Fields) Fields {
			fields["trace_id"] = "trace-1"
			return fields
		},
	}).(*Loggers)

	prepared := l.prepare(context.Background(), Any("k", "v"))

	if prepared["module"] != "test" {
		t.Fatalf("module missing, got: %#v", prepared["module"])
	}
	if prepared["k"] != "v" {
		t.Fatalf("field missing, got: %#v", prepared["k"])
	}
	if prepared["trace_id"] != "trace-1" {
		t.Fatalf("ToFieldsFunc not applied, got: %#v", prepared["trace_id"])
	}
}
