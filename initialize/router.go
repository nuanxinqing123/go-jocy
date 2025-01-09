package initialize

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"go-jocy/config"
	"go-jocy/internal/middleware"
	"go-jocy/internal/router"
)

func Routers() *gin.Engine {
	fmt.Println(" ")
	if config.GinConfig.App.Mode == "debug" {
		fmt.Println("运行模式: Debug模式")
		gin.SetMode(gin.DebugMode)
	} else {
		fmt.Println("运行模式: Release模式")
		gin.SetMode(gin.ReleaseMode)
	}
	fmt.Println("监听端口: " + strconv.Itoa(config.GinConfig.App.Port))
	fmt.Println(" ")

	Router := gin.New()
	Router.Use(middleware.Logger())
	Router.Use(middleware.Recovery())

	// 允许跨域
	Router.Use(cors.New(middleware.CorsConfig))

	// (可选项)
	// PID 限流基于实例的 CPU 使用率，通过拒绝一定比例的流量, 将实例的 CPU 使用率稳定在设定的阈值上。
	// 地址: https://github.com/bytedance/pid_limits
	// Router.Use(adaptive.PlatoMiddlewareGinDefault(0.8))

	// 存活检测
	Router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	ApiGroup := Router.Group("/app")
	router.InitRouterApp(ApiGroup)

	return Router
}
