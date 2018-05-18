package conv

import "strconv"

func intToStr(i int) string {
	return strconv.Itoa(i)
}

func int32ToStr(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}
