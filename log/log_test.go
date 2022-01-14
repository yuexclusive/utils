package log

import (
	"testing"

	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	Panic("error", zap.String("name", "this is a test log"))
}
