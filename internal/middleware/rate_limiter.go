package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type RequestInfo struct {
	LastAccessTime time.Time // 上次访问时间
	RequestNum     int       // 请求计数
}

var (
	requestInfoMap = make(map[string]*RequestInfo) // IP到请求信息的映射
	mutex          = &sync.Mutex{}                 // 用于保护requestInfoMap的互斥锁
	maxRequests    = 180                           // 允许的最大请求数
	timeWindow     = 1 * time.Minute               // 时间窗口
)

// RateLimitMiddleware 限流
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		mutex.Lock()
		defer mutex.Unlock()
		// 检查IP是否在map中
		info, exists := requestInfoMap[ip]
		// 如果IP不存在，初始化并添加到map中
		if !exists {
			requestInfoMap[ip] = &RequestInfo{LastAccessTime: time.Now(), RequestNum: 1}
			return
		}
		// 如果IP存在，检查时间窗口
		if time.Since(info.LastAccessTime) > timeWindow {
			// 如果超过时间窗口，重置请求计数
			info.RequestNum = 1
			info.LastAccessTime = time.Now()
			return
		}
		info.RequestNum++ // 如果在时间窗口内，增加请求计数
		// 如果请求计数超过限制，禁止访问
		if info.RequestNum > maxRequests {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			c.Abort()
			return
		}
		// 更新最后访问时间
		info.LastAccessTime = time.Now()
		c.Next()
	}
}
