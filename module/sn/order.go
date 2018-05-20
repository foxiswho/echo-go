package sn

import "github.com/foxiswho/echo-go/util/sn"

func MakeOrderNo() string {
	return sn.MakeYearDaysRand(12)
}
