package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var cfg *Config
var log *zap.Logger
var sugaredLogger *zap.SugaredLogger

type Config struct {
	Level      zapcore.Level `json:"level"`
	Filename   string        `json:"filename"`
	MaxSize    int           `json:"max_size"`
	MaxAge     int           `json:"max_age"`
	MaxBackups int           `json:"max_backups"`
}

func InitConfig() {
	var level = zapcore.InfoLevel
	var filename = ""
	var maxSize = 300
	var maxAge = 30
	var maxBackups = 7

	filename = viper.GetString("log.filename")
	if ls := viper.GetString("log.level"); ls != "" {
		var l = zapcore.InfoLevel
		err := l.UnmarshalText([]byte(ls))
		if err == nil {
			level = l
		}
	}

	if ms := viper.GetInt("log.max_size"); ms > 0 {
		maxSize = ms
	}

	if ma := viper.GetInt("log.max_age"); ma > 0 {
		maxAge = ma
	}

	if mb := viper.GetInt("log.max_age"); mb > 0 {
		maxBackups = mb
	}

	cfg = &Config{
		Level:      level,
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
	}
}

// Init 初始化 Logger
func Init() {
	InitConfig()
	writeSyncer := getLogWriter(
		cfg.Level,
		cfg.Filename,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge,
	)
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writeSyncer, cfg.Level)

	log = zap.New(core, zap.AddCaller())
	sugaredLogger = log.Sugar()
	// 替换zap库中全局的logger
	zap.ReplaceGlobals(log)
	return
}

func GetLogger() *zap.Logger {
	return log
}

func GetSugaredLogger() *zap.SugaredLogger {
	return sugaredLogger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(level zapcore.Level, filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	if filename != "" {
		return zapcore.AddSync(zapcore.AddSync(os.Stdout))
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	if level <= zapcore.DebugLevel {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))
	}

	return zapcore.AddSync(lumberJackLogger)

}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
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
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: err check
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
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
