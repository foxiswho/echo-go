package file

import (
	"encoding/base64"
	"github.com/foxiswho/shop-go/util/conv"
	"github.com/foxiswho/shop-go/util/crypt"
	"github.com/foxiswho/shop-go/util"
	"github.com/foxiswho/shop-go/conf"
)

//令牌生成
//@maps 令牌数组
//
func TokeMake(maps map[string]interface{}) (string, error) {
	s, err := conv.ObjToJson(maps)
	if err != nil {
		return "", util.NewError("序列化失败：" + err.Error())
	}
	key := []byte(conf.Conf.Secret.UploadAesKey)
	result, err := crypt.AesEncrypt([]byte(s), key)
	if err != nil {
		return "", util.NewError("加密失败：" + err.Error())
	}
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	return b64.EncodeToString(result), nil
}

//令牌解密
//@str 加密的字符串
func TokenDeCode(str string) (map[string]interface{}, error) {
	if len(str) < 1 {
		return nil, util.NewError("字符串 不能为空")
	}
	key := []byte(conf.Conf.Secret.UploadAesKey)
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	byt, err := b64.DecodeString(str)
	if err != nil {
		return nil, util.NewError("base64解码失败：" + err.Error())
	}
	origData, err := crypt.AesDecrypt(byt, key)
	if err != nil {
		return nil, util.NewError("解密失败：" + err.Error())
	}
	maps, err := conv.StrToMap(string(origData))
	if err != nil {
		return nil, util.NewError("转换为map失败：" + err.Error())
	}
	return maps, nil
}
