package conv

import (
	"encoding/json"
	"github.com/foxiswho/shop-go/util"
)

//JSON格式数据转换为map
func StrToMap(str string) (mapObj map[string]interface{}, err error) {
	// 结构体转json
	if str == "" {
		return nil, util.NewError("字符串为空不能转换")
	}
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		return nil, err
	}
	return result, nil
}
