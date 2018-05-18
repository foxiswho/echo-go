package conv

import (
	"encoding/json"
	"fmt"
	"github.com/foxiswho/echo-go/util"
)

//obj 反序列化为字符串
func ObjToJson(v interface{}) (string, error) {
	str, err := json.Marshal(v)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return "", util.NewError("序列化失败:" + err.Error())
	}
	return string(str), nil
}
