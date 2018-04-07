package common

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"math/rand"
)

func TakeMd5(pass string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(pass))
	md5Data := hex.EncodeToString(md5Ctx.Sum(nil))
	return md5Data
}

// return len=8  salt
func GetRandomSalt() string {
	return GetRandomString(8)
}

//生成随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

