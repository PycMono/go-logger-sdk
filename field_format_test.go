package logsdk

import (
	"context"
	"testing"
	"time"

	goerrors "github.com/go-errors/errors"
)

type samplePayload struct {
	Name string `json:"name"`
}

func TestFieldsFormatReturnsNormalizedValues(t *testing.T) {
	Info(context.TODO(), "测试", Any("1", "3"))

	now := time.Date(2026, 4, 7, 10, 11, 12, 0, time.UTC)
	in := Fields{
		"bytes":    []byte("abc"),
		"duration": 2 * time.Second,
		"time":     now,
		"num":      int64(7),
		"str":      "ok",
		"struct":   samplePayload{Name: "demo"},
		"nil":      nil,
	}

	out := in.format()

	if out["bytes"] != "abc" {
		t.Fatalf("bytes not formatted, got: %#v", out["bytes"])
	}
	if out["duration"] != "2s" {
		t.Fatalf("duration not formatted, got: %#v", out["duration"])
	}
	if out["time"] != now.Format(time.RFC3339) {
		t.Fatalf("time not formatted, got: %#v", out["time"])
	}
	if out["num"] != int64(7) {
		t.Fatalf("numeric value changed, got: %#v", out["num"])
	}
	if out["str"] != "ok" {
		t.Fatalf("string value changed, got: %#v", out["str"])
	}
	if out["struct"] != `{"name":"demo"}` {
		t.Fatalf("struct not marshaled, got: %#v", out["struct"])
	}
	if _, ok := out["nil"]; !ok {
		t.Fatalf("nil key missing from output")
	}
}

func TestFieldsFormatAddsErrorStackForGoErrors(t *testing.T) {
	out := Fields{"error": goerrors.New("boom")}.format()

	if _, ok := out["errorsStack"]; !ok {
		t.Fatalf("expected errorsStack for standard error, got: %#v", out)
	}
}
