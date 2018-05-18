package model

import (
	"github.com/foxiswho/shop-go/module/log"
)

type Logger struct {
}

// Print format & print log
func (logger Logger) Print(values ...interface{}) {
	// @TODO
	// 日志格式化解析
	log.Debugf("orm log:%v \n", values)
}
