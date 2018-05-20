package sn

import (
	"time"
	"github.com/foxiswho/echo-go/util/datetime"
	"github.com/foxiswho/echo-go/util/str"
	"strings"
	"math"
)

//生成单号
//06123xxxxx
//sum 最少10
func MakeYearDaysRand(sum int) string {
	//年
	strs := time.Now().Format("06")
	//一年中的第几天
	days := string(datetime.GetDaysInYearByThisYear())
	count := len(days)
	if count < 3 {
		//重复字符0
		days = strings.Repeat("0", 3-count) + days
	}
	//组合
	strs += days
	//剩余随机数
	sum = sum - 5
	if sum < 1 {
		sum = 5
	}
	//0~9999999的随机数
	ran := str.GetRand()
	result := string(ran.Intn(int(math.Pow(0, float64(sum+1)) - 1)))
	count = len(result)
	if count < sum {
		//重复字符0
		result = strings.Repeat("0", sum-count) + result
	}
	//组合
	strs += result
	return strs
}
