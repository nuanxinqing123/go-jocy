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
