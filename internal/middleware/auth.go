package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"go-jocy/utils"
)

// CheckIfWithinUTC8 checks if the given time is within ±10 seconds of the current UTC+8 time
func CheckIfWithinUTC8(inputTime time.Time) bool {
	// 定义 UTC+8 时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return false
	}

	// 当前 UTC+8 时间
	nowUTC8 := time.Now().In(location)

	// 计算时间差（秒）
	timeDiff := nowUTC8.Sub(inputTime).Seconds()

	// 检查是否在 ±10 秒范围内
	return timeDiff >= -10 && timeDiff <= 10
}

// Auth 请求签名校验
func Auth() gin.HandlerFunc {
	/*
		过程：jocy&时间戳&随机字符.反转随机字符
		示例：jocy&1733128303&abcd.dcba
	*/
	return func(c *gin.Context) {
		// Header中获取签名
		ts := c.Request.Header.Get("t")
		sign := c.Request.Header.Get("s")

		// 判断签名是否存在
		if ts == "" || sign == "" {
			c.JSON(403, gin.H{
				"msg": "禁止未经授权访问",
			})
			c.Abort()
			return
		}

		// String时间戳转换为Unix时间戳
		t, err := time.Parse("2006-01-02 15:04:05", ts)
		if err != nil {
			c.JSON(403, gin.H{
				"msg": "禁止未经授权访问",
			})
			c.Abort()
			return
		}

		// 校验时间戳[是否在UTC+8时区的+-10秒以内]
		if CheckIfWithinUTC8(t) == false {
			c.JSON(403, gin.H{
				"msg": "禁止未经授权访问",
			})
			c.Abort()
			return
		}

		// 从Sign中使用[.]分割
		signs := strings.Split(sign, ".")
		if len(signs) < 2 {
			c.JSON(403, gin.H{
				"msg": "禁止未经授权访问",
			})
			c.Abort()
			return
		}
		s := signs[0]  // 签名
		rs := signs[1] // 反转随机字符

		// 生成签名
		sn := fmt.Sprintf("jocy&%s&%d&%s", s, t.Unix(), utils.ReverseString(rs))

		// 判断签名是否一致
		if utils.MD5Encryption(sn) != s {
			c.JSON(403, gin.H{
				"msg": "禁止未经授权访问",
			})
			c.Abort()
			return
		}

		// 认证通过
		c.Next()
	}
}
