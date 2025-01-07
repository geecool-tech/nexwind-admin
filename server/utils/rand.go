package utils

import (
	"math/rand"
	"time"
)

// init函数设置随机数生成器的种子为当前时间，提高随机性。
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateShortUUID 生成一个8位的随机字符串作为短UUID。
func GenerateShortUUID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, charset[rand.Intn(len(charset))])
	}
	return string(result)
}
