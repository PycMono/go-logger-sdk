package logdk

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
)

// CallerKey holds the caller stack field
const CallerKey = "caller"

const (
	stackJump       = 4
	fieldsStackJump = 6
)

// Caller decorates log entries with function name and code line number
type Caller struct {
	Formatter logrus.Formatter
}

// Format the current log entry by adding the function name and line number of the caller.
func (f *Caller) Format(entry *logrus.Entry) ([]byte, error) {
	function, file, line := f.getCurrentPosition(entry)

	packageEnd := strings.LastIndex(function, ".")
	functionName := function[packageEnd+1:]

	caller := fmt.Sprintf("%s[%s:%d]", functionName, filepath.Base(file), line)
	data := logrus.Fields{CallerKey: caller} // 设置caller字段
	for k, v := range entry.Data {
		data[k] = v
	}
	entry.Data = data

	return f.Formatter.Format(entry)
}

func (f *Caller) getCurrentPosition(entry *logrus.Entry) (string, string, int) {
	skip := stackJump
	if len(entry.Data) == 0 {
		skip = fieldsStackJump
	}
start:
	pc, file, line, _ := runtime.Caller(skip)
	function := runtime.FuncForPC(pc).Name()
	if strings.LastIndex(function, "sirupsen/logrus.") >= 0 ||
		strings.LastIndex(function, "log/logrusx.") >= 0 {
		skip++
		goto start
	}
	return function, file, line
}
