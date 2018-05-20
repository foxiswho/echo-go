package datetime

import "time"

//Unix时间戳
func UnixTime() int64 {
	return time.Now().Unix()
}

func UnixTimeFormat(i int64, str string) string {
	return time.Unix(i, 0).Format(str)
}

func UnixTimeFormatDateTime(i int64) string {
	return time.Unix(i, 0).Format(Y_M_D_H_I_S)
}
