package middleware

import (
	"Go-Mall/app/client/conf"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	hertzobslogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// RegisterLogMiddleware 注册日志中间件
func RegisterLogMiddleware() {
	// 初始化日志
	logger := hertzobslogrus.NewLogger(hertzobslogrus.WithLogger(hertzobslogrus.NewLogger().Logger()))
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())

	var flushInterval time.Duration
	// 生产环境刷新间隔设置为1分钟
	if os.Getenv("GO_ENV") == "online" {
		flushInterval = time.Minute
	} else {
		flushInterval = time.Second
	}

	// 设置异步写入
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: flushInterval,
	}

	// 设置日志输出
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
}
