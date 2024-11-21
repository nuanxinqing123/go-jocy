package utils

import (
	"math/rand"
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
