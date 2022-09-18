package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// 获取md5
func Md5Encode(str string) string {
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func GenPassword(pwd string, sale string) string {
	return Md5Encode(fmt.Sprintf("%s-%s", pwd, sale))
}
