package gobtphelper

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	hashBytes := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
