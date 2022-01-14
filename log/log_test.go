package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	_driver.Logger().Error("test")
}
