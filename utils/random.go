package utils

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// RandomChoice 从 []string 中随机选择一个元素
func RandomChoice(options []string) string {
	if len(options) == 0 {
		return "" // 返回空字符串，表示输入为空
	}

	// 设置随机种子
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// 随机索引
	index := rand.Intn(len(options))

	return options[index]
}

// RandomString 生成指定长度的随机字符串
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// ReverseString 字符串反转
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// RandomGetElements 从切片中随机获取指定数量的元素
func RandomGetElements(urls []string) string {
	// 检查URL列表是否为空
	if len(urls) == 0 {
		return ""
	}
	// 初始化随机数种子
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成随机索引
	randomIndex := rand.Intn(len(urls))
	return urls[randomIndex]
}

// BindUserToUrl 绑定用户ID到URL地址
func BindUserToUrl(userId string, urls []string) (string, error) {
	// 检查URL列表是否为空
	if len(urls) == 0 {
		return "", errors.New("URL列表不能为空")
	}
	// 转化为整数
	id, err := strconv.Atoi(userId)
	if err != nil {
		return "", errors.New("用户ID必须是一个有效的数字字符串")
	}

	// 计算索引，取绝对值防止负数
	index := int(math.Abs(float64(id))) % len(urls)
	return urls[index], nil
}
