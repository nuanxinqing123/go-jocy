package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"go-jocy/config"
	"go-jocy/initialize"
	"go-jocy/internal/cron"
)

func main() {
	// 初始化配置
	config.GinVP = initialize.Viper()

	// 初始化日志
	config.GinLOG = initialize.Zap()

	// 启动定时任务
	if err := cron.InitTask(); err != nil {
		fmt.Printf("定时任务初始化失败, err:%v\n", err)
		return
	}
	zap.L().Debug("定时任务初始化成功...")

	// 初始化路由
	router := initialize.Routers()
	if router == nil {
		fmt.Println("初始化路由失败...")
		return
	}

	// 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GinConfig.App.Port),
		Handler: router,
	}

	// 启动
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Listten: %s\n", err)
		}
	}()

	// 等待终端信号来优雅关闭服务器，为关闭服务器设置10秒超时
	quit := make(chan os.Signal, 1) // 创建一个接受信号的通道

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞此处，当接受到上述两种信号时，才继续往下执行
	config.GinLOG.Info("Service ready to shut down")

	// 创建10秒超时的Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 10秒内优雅关闭服务（将未处理完成的请求处理完再关闭服务），超过10秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		config.GinLOG.Fatal("Service timed out has been shut down: ", zap.Error(err))
	}

	config.GinLOG.Info("Service has been shut down")
}
