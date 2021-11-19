package log

import (
	"time"

	"github.com/gin-gonic/gin"
)

var (
	_driver *Zap
)

func init() {
	if _driver == nil {
		_driver = NewZap()
	}
}

// Debug logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func Debug(msg string, keysAndValues ...interface{}) {
	_driver.Sugar().Debugw(msg, keysAndValues...)
}

// Info logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Info(msg string, keysAndValues ...interface{}) {
	_driver.Sugar().Infow(msg, keysAndValues...)
}

// Warn logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warn(msg string, keysAndValues ...interface{}) {
	_driver.Sugar().Warnw(msg, keysAndValues...)
}

// Error logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Error(msg string, keysAndValues ...interface{}) {
	_driver.Sugar().Errorw(msg, keysAndValues...)
}

// DPanic logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanic(msg string, keysAndValues ...interface{}) {
	_driver.Sugar().DPanicw(msg, keysAndValues...)
}

// Panic logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panic(msg string, keysAndValues ...interface{}) {
	_driver.Sugar().Panicw(msg, keysAndValues...)
}

// Fatal logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatal(msg string, keysAndValues ...interface{}) {
	_driver.Sugar().Fatalw(msg, keysAndValues...)
}

// GinUseZap GinUseZap
func GinUseZap(engine *gin.Engine) {
	engine.Use(_driver.Ginzap(time.RFC3339, true))
	engine.Use(_driver.RecoveryWithZap(true))
}
