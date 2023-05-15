package utils

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// GenerateUsername 生成随机用户名
func GenerateUsername(perfix string, length int) string {
	rand.Seed(time.Now().UnixNano())
	min := int(math.Pow10(length - 1))
	max := int(math.Pow10(length)) - 1
	username := fmt.Sprintf("%s%d", perfix, rand.Intn(max-min+1)+min)
	return username
}

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateInviteCode 生成随机推荐码
func GenerateInviteCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, length)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}

// GenerateOrderNumber 生成订单号函数
func GenerateOrderNumber(prefix string) string {

	// 将当前时间戳转换成字符串
	timeStr := time.Now().Format("20060102150405")
	// 生成5位随机数
	rand.Seed(time.Now().UnixNano())
	min := 10000
	max := 99999
	r := rand.Intn(max-min+1) + min
	// 将随机数和时间戳拼接成字符串
	orderNumber := fmt.Sprintf("%s%s%d", prefix, timeStr, r)
	return orderNumber
}
