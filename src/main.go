package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"hole/src/config"
	"hole/src/config/logger"
	"hole/src/config/mysql"
	"hole/src/config/redis"
	"hole/src/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Go Web开发较通用的脚手架模板

func main() {
	// 1. 加载配置
	config.InitConfigFile()
	// 2. 初始化日志
	logger.InitLogger()
	logger.GetLogger().Info("logger init success...")
	// 3. 初始化 mysql
	mysql.InitMysql()
	// 4. 初始化Redis连接
	redis.InitRedis()
	// 5. 注册路由
	r := routers.Setup()
	// 6. 启动服务（优雅关机）
	logger.GetLogger().Debug("listener port ...", zap.Int("port", config.GetPort()))
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetPort()),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.GetSugaredLogger().Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}