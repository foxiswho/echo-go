package main

import (
	"flag"
	"fmt"

	"github.com/foxiswho/echo-go/module/log"
	"github.com/foxiswho/echo-go/router"
)

var helpInfo = "help\n  -h      		帮助\n  -c conf/conf.toml	配置文件路径，默认conf/conf.toml\n"
var cmdConf = flag.Bool("c", false, "配置文件路径")
var cmdHelp = flag.Bool("h", false, "帮助")

func main() {
	confFilePath := "" // 默认conf/conf.toml

	//解析命令行标志
	flag.Parse() // Scans the arg list and sets up flags
	if *cmdConf {
		if flag.NArg() == 1 {
			confFilePath = flag.Arg(0)
			log.Debugf("run with conf:%s", confFilePath)
		} else {
			fmt.Printf("-c 参数错误\n" + helpInfo)
			return
		}
	} else if *cmdHelp {
		fmt.Printf(helpInfo)
		return
	}

	// 子域名部署
	router.RunSubdomains(confFilePath)
}
