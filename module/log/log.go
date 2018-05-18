package log

import (
	l "github.com/labstack/gommon/log"
)

type Logger struct {
	*l.Logger
}

var (
	// 自定义全局log
	// 使用gommon/log的global由于限制3层调用栈获取不到log的准确file路径
	global        = l.New("log")
	defaultHeader = `{"time":"${time_rfc3339}","level":"${level}","prefix":"${prefix}",` +
		`"file":"${long_file}","line":"${line}"}`
)

func init() {
	// 全局log属性设置
	l.SetHeader(defaultHeader)
	l.SetLevel(l.DEBUG)

	// 自定义全局logs属性设置
	global.SetHeader(defaultHeader)
	global.SetLevel(l.DEBUG)
}

func Debugf(format string, values ...interface{}) {
	global.Debugf(format, values...)
}
