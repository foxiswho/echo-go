#!/bin/bash

OS="local"
while getopts "lh:" arg
do
    case $arg in

        h)  #help带all参数仅用于测试脚本，并备忘
            echo "-h [all] help"
            case $OPTARG in
                all)
                    echo "-l build linux bin, default local"
                    ;;
            esac
            exit
            ;;
        l)
            OS="linux"
            ;;
        ?)
            echo "unkonw argument"
            echo "-h all help"
            echo "-l build linux bin, default local"
            exit 1
            ;;
    esac
done

# Go Path
# CURDIR=`pwd`
# OLDGOPATH="$GOPATH"
# export GOPATH="$OLDGOPATH:$CURDIR"

LogPrefix=">>>>"

# 打包前检测Bindata是否开启

echo -e "$LogPrefix `date +"%H:%M:%S"` \033[42;37m start \033[0m"

echo "$LogPrefix `date +"%H:%M:%S"` assets bindata"
go-bindata -ignore=\\.DS_Store -ignore=assets.go -pkg="assets" -o assets/assets.go assets/...

echo "$LogPrefix `date +"%H:%M:%S"` template bindata"
go-bindata -ignore=\\.DS_Store -ignore=template.go -pkg="template" -o template/template.go template/...

# echo "$LogPrefix `date +"%H:%M:%S"` src package"
# gofmt -w src/

# 交叉编译
case  $OS  in   
    linux)  
        # Linux
        echo "$LogPrefix `date +"%H:%M:%S"` build linux bin"
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o echo-web ./echo.go
        ;;  
    *) 
        # 本机
        echo "$LogPrefix `date +"%H:%M:%S"` build local bin"
        go build -a -installsuffix cgo -ldflags '-w' -o echo-web ./echo.go
        ;;
esac 

echo -e "$LogPrefix `date +"%H:%M:%S"` \033[42;37m finished \033[0m"