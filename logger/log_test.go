package logger

import (
	"testing"
)

func TestLog(t *testing.T) {
	Info("test info")
	Error("test error")
}
