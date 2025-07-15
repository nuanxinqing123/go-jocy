package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"

	"go-jocy/config"
)

type Request struct {
	client    *resty.Client
	AuthToken string
}

func New(AuthToken, AuthIP string) *Request {
	client := resty.New()

	// 设置DeBug
	if config.GinConfig.App.Mode == "debug" {
		client.SetDebug(true)
	}

	// 设置错误重试
	client.SetRetryCount(5)
	client.SetRetryWaitTime(1 * time.Second)

	// 设置客户端IP
	if AuthIP != "" {
		client.SetHeaderVerbatim("X-Forwarded-For", AuthIP)
		client.SetHeaderVerbatim("X-Real-IP", AuthIP)
		client.SetHeaderVerbatim("True-Client-IP", AuthIP)
		client.SetHeaderVerbatim("Client-IP", AuthIP)
	}

	// 计算签名
	ts := time.Now().UnixMilli()
	signText := fmt.Sprintf("2.4.6.5-%s-Android-1.5.7.0-3b11fa670075426f", strconv.FormatInt(ts, 10))
	sign, _ := AesEncryptionBase64(signText, "ziISjqkXPsGUMRNGyWigxDGtJbfTdcGv", "WonrnVkxeIxDcFbv")

	// 设置请求头
	client.SetHeaderVerbatim("User-Agent", "Dart/2.17 (dart:io)")
	client.SetHeaderVerbatim("Accept-Encoding", "gzip")
	client.SetHeaderVerbatim("x-version", "2020-09-17")
	client.SetHeaderVerbatim("appid", "4150439554430529")
	client.SetHeaderVerbatim("ts", strconv.FormatInt(ts, 10))
	client.SetHeaderVerbatim("authentication", sign)
	client.SetHeaderVerbatim("tcs", "2")

	if AuthToken != "" {
		client.SetHeaderVerbatim("x-token", AuthToken)
	}
	return &Request{
		client:    client,
		AuthToken: AuthToken,
	}
}

func (r *Request) Get(url string, params map[string]string) (*resty.Response, error) {
	config.GinLOG.Debug("GET: " + url)
	config.GinLOG.Debug(fmt.Sprintf("params: %v", params))

	return r.client.R().
		SetQueryParams(params).
		Get(url)
}

func (r *Request) Post(url string, body any) (*resty.Response, error) {
	config.GinLOG.Debug("POST: " + url)
	config.GinLOG.Debug(fmt.Sprintf("body: %v", body))

	return r.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)
}

func (r *Request) Delete(url string, body any) (*resty.Response, error) {
	config.GinLOG.Debug("DELETE: " + url)
	config.GinLOG.Debug(fmt.Sprintf("body: %v", body))

	return r.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Delete(url)
}
