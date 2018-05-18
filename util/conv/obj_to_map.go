package conv

import (
	"encoding/json"
)

// 函　数：Obj2map
// 概　要：
// 参　数：
//      obj: 传入Obj
// 返回值：
//      mapObj: map对象
//      err: 错误
func ObjToMap(obj interface{}) (mapObj map[string]interface{}, err error) {
	// 结构体转json
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}
