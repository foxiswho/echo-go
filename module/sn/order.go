package sn

import (
	"github.com/foxiswho/echo-go/util/str"
)

func MakeOrderNo() string {
	return str.MakeYearDaysRand(12)
}
