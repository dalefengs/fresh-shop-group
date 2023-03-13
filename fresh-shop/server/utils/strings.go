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
