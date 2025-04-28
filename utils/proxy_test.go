package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

// TestHttpProxy 测试HTTP代理是否可用
func TestHttpProxy(t *testing.T) {
	// 代理地址
	// proxy := "http://47.120.43.237:8686"
	proxy := "socks5h://8.134.14.215:34567"

	testURL := "http://httpbin.org/ip"
	timeout := 5 * time.Second

	// 创建一个新的resty客户端
	client := resty.New()

	// 设置代理
	client.SetProxy(proxy)

	// 设置超时
	client.SetTimeout(timeout)

	// 发送请求
	resp, err := client.R().Get(testURL)

	if err != nil {
		t.Logf("代理不可用，错误信息：%v", err)
		return
	}

	// 检查状态码
	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		t.Logf("代理可用，返回内容：%s", resp.String())
	} else {
		t.Logf("代理不可用，状态码：%d，返回内容：%s", resp.StatusCode(), resp.String())
	}
}

// TestMultipleProxies 测试多个代理
func TestMultipleProxies(t *testing.T) {
	proxies := []string{
		"http://47.120.43.237:8686",
		"socks5h://8.134.14.215:34567",
		// 可以添加更多代理地址进行测试
	}

	testURL := "http://httpbin.org/ip"
	timeout := 5 * time.Second

	for _, proxy := range proxies {
		t.Run(fmt.Sprintf("测试代理: %s", proxy), func(t *testing.T) {
			// 创建一个新的resty客户端
			client := resty.New()

			// 设置代理
			client.SetProxy(proxy)

			// 设置超时
			client.SetTimeout(timeout)

			// 发送请求
			resp, err := client.R().Get(testURL)

			if err != nil {
				t.Logf("代理不可用，错误信息：%v", err)
				return
			}

			// 检查状态码
			if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
				t.Logf("代理可用，返回内容：%s", resp.String())
			} else {
				t.Logf("代理不可用，状态码：%d，返回内容：%s", resp.StatusCode(), resp.String())
			}
		})
	}
}
