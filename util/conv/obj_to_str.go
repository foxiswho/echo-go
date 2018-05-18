package conv

func ObjToStr(i interface{}) string {
	if i == nil {
		return ""
	}
	return i.(string)
}
