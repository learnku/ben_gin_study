package utils_password

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5 加密（小写）
func Md5Encode(data string) string {
	h := md5.New()
	// 需要加密的字符串为
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

// Md5 加密（全大写）
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}
