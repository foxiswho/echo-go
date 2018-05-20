package str

import (
	"math/rand"
	"time"
)

// randseed
func GetRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// rand salt
func RandSalt() string {
	var salt = ""
	for i := 0; i < 4; i++ {
		ran := GetRand()
		salt += string(SALT[ran.Intn(len(SALT))])
	}
	return salt
}

const (
	SALT = "$^*#,.><)(_+f*m"
)

//生成 数字和小写字母
//https://blog.csdn.net/qq948993066/article/details/77368971
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//生成随机字符串
func GetRandomString2(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
