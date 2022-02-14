package logger

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"runtime/debug"
	"strings"
	"time"

	"config"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// DevelopMode dev mod config
	Production = "production"
)

// Zap Zap
type Zap struct {
	zapLogger *zap.Logger
	zapConfig *config.Log
}

// Level Level
type Level int8

// Enabled implement LevelEnabler
func (l Level) Enabled(in zapcore.Level) bool {
	return l == Level(in)
}

// NewZap NewZap
func NewZap() *Zap {
	z := &Zap{}
	z.initConfig()
	z.initLogger()
	return z
}

func (z *Zap) initConfig() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("read config for log failed: ", r)
		}
	}()
	cfg := config.Init[config.Config]("config.toml").GetConfig()
	z.zapConfig = cfg.Log
}

// new logger
func (z *Zap) initLogger() {
	cores := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(z.getConsoleEncoderConfig()), os.Stdout, zapcore.DebugLevel),
		z.newCore(zapcore.DebugLevel),
		z.newCore(zapcore.InfoLevel),
		z.newCore(zapcore.WarnLevel),
		z.newCore(zapcore.ErrorLevel),
		z.newCore(zapcore.DPanicLevel),
		z.newCore(zapcore.PanicLevel),
		z.newCore(zapcore.FatalLevel),
	)

	options := []zap.Option{
		// zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	}

	if z.zapConfig == nil || strings.ToLower(strings.TrimSpace(z.zapConfig.Mode)) != Production {
		options = append(options, zap.Development())
	}

	logger := zap.New(cores, options...)

	z.zapLogger = logger
}

// newCore newCore
func (z *Zap) newCore(level zapcore.Level) zapcore.Core {
	return zapcore.NewCore(zapcore.NewJSONEncoder(z.getESEncoderConfig()), z.getLogWriter(level.String()), Level(level))
}

// getLogWriter 写入文件
func (z *Zap) getLogWriter(level string) zapcore.WriteSyncer {
	filePath := "./log"
	if z.zapConfig != nil {
		filePath = z.zapConfig.Path
	}
	hook, err := rotatelogs.New(
		// path.Join(filePath, "%Y-%m-%d_"+fmt.Sprintf("%s.log", level)),
		path.Join(filePath, "%Y-%m-%d-%H_"+fmt.Sprintf("%s.log", level)),
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(time.Hour*12),
	)

	if err != nil {
		log.Fatalf("log getLogWriter rotatelogs.New: %s", err)
	}
	return zapcore.AddSync(hook)
}

func (z *Zap) getConsoleEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder    // 时间格式 RFC3339
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 级别大写+颜色显示
	return encoderConfig
}

func (z *Zap) getESEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder // 时间格式 RFC3339
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder   // 级别大写+颜色显示
	return encoderConfig
}

// Logger Logger
func (z *Zap) Logger() *zap.Logger {
	return z.zapLogger
}

// Sugar Sugar
func (z *Zap) Sugar() *zap.SugaredLogger {
	return z.zapLogger.Sugar()
}

// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//   1. A time package format string (e.g. time.RFC3339).
//   2. A boolean stating whether to use UTC time zone or local.
func (z *Zap) Ginzap(timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				z.Logger().Error(e)
			}
		} else {
			z.Logger().Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("time", end.Format(timeFormat)),
				zap.Duration("latency", latency),
			)
		}
	}
}

// RecoveryWithZap returns a gin.HandlerFunc (middleware)
// that recovers from any panics and logs requests using uber-go/zap.
// All errors are logged using zap.Error().
// stack means whether output the stack info.
// The stack info is easy to find where the error occurs but the stack info is too large.
func (z *Zap) RecoveryWithZap(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					z.Logger().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					z.Logger().Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					z.Logger().Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

// GinUseZap GinUseZap
func GinUseZap(engine *gin.Engine) {
	engine.Use(_driver.Ginzap(time.RFC3339, true))
	engine.Use(_driver.RecoveryWithZap(true))
}
