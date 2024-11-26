package initialize

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	assetfs "github.com/elazarl/go-bindata-assetfs"

	"go-jocy/config"
	"go-jocy/internal/middleware"
	"go-jocy/internal/router"
	"go-jocy/web/bindata"
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

	// 获取CDN地址
	Router.Use(func(c *gin.Context) {
		// 优先从上游Header中获取客户端IP
		clientIP := c.Request.Header.Get("X-Real-IP")
		if clientIP == "" {
			clientIP = c.Request.Header.Get("X-Forwarded-For")
		}
		if clientIP == "" {
			clientIP = c.Request.Header.Get("True-Client-IP")
		}
		if clientIP == "" {
			clientIP = c.Request.Header.Get("Client-IP")
		}
		if clientIP == "" {
			clientIP = c.Request.RemoteAddr
		}
		c.Set("x-client-ip", clientIP)
		c.Next()
	})

	// (可选项)
	// PID 限流基于实例的 CPU 使用率，通过拒绝一定比例的流量, 将实例的 CPU 使用率稳定在设定的阈值上。
	// 地址: https://github.com/bytedance/pid_limits
	// Router.Use(adaptive.PlatoMiddlewareGinDefault(0.8))

	// 前端静态文件
	{
		// 加载模板文件
		t, err := loadTemplate()
		if err != nil {
			panic(err)
		}
		Router.SetHTMLTemplate(t)

		// 加载静态文件
		fs := &assetfs.AssetFS{
			Asset:     bindata.Asset,
			AssetDir:  bindata.AssetDir,
			AssetInfo: nil,
			Prefix:    "assets",
		}
		Router.StaticFS("/assets", fs)

		Router.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", nil)
		})
	}

	// 存活检测
	Router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	ApiGroup := Router.Group("/app")
	router.InitRouterApp(ApiGroup)

	return Router
}

// loadTemplate 加载模板文件
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for _, name := range bindata.AssetNames() {
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		asset, err := bindata.Asset(name)
		if err != nil {
			continue
		}
		name = strings.Replace(name, "assets/", "", 1)
		t, err = t.New(name).Parse(string(asset))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
