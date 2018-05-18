package conv

import "strconv"

//onj变成数字
func ObjToInt(i interface{}) (int, error) {
	n := 0
	switch i.(type) {
	case int:
		n = i.(int)
	case int32:
		n = int(i.(int32))
	case int64:
		n = int(i.(int64))
	case float32:
		n = int(i.(float32))
	case float64:
		n = int(i.(float64))
	case string:
		var err error
		n, err = strconv.Atoi(i.(string))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}
