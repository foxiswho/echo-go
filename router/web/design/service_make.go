package design

import (
	"github.com/foxiswho/echo-go/router/base"
	"github.com/foxiswho/echo-go/module/db"
	"fmt"
	"text/template"
	"os"
	"github.com/foxiswho/echo-go/conf"
	"net/http"
	"github.com/foxiswho/echo-go/util/str"
	"strings"
)

func ServiceMakeHandler(c *base.BaseContext) error {

	sql := "show tables"
	result, err := db.DB().Engine.QueryString(sql)
	fmt.Println("err", err)
	fmt.Println("result", result)
	if err != nil {
		fmt.Println("err", err)
	} else {
		template_file := "./template/design/make/service.go.tpl" //模版
		service_path := "./service/base"                         //生成到目录
		table_field := "Tables_in_" + conf.Conf.DB.Name          //获取字段名称
		prefix := ""                                             //前缀
		make_file_suffix := ""                                   //生成文件后缀
		for i, val := range result {
			field := val[table_field]
			field = strings.Replace(field, prefix, "", -1) //替换掉表前缀
			fmt.Println("result i val=>", i, val)
			fmt.Println("result field=>", field)
			tmpl, err := template.ParseFiles(template_file)
			fmt.Println("template err", err)
			fmt.Println("template err", tmpl)
			//field = "admin_menu"
			data := make(map[string]interface{})
			data["tables"] = field
			//data["tables_first"] = gonicCasedName(field)
			data["tables_camel_case"] = str.BigCamelCase(field)           //驼峰命名法 大写
			data["tables_big_camel_case"] = data["tables_camel_case"]     //驼峰命名法 大写
			data["tables_little_camel_case"] = str.LittleCamelCase(field) //驼峰命名法 小写
			fmt.Println("data=>", data)
			fmt.Println("data=>", field)
			//
			//
			//
			err = os.MkdirAll(service_path, os.ModePerm)
			if err != nil {
				fmt.Println("%s", err)
			} else {
				fmt.Println("Create Directory OK! ", service_path)
			}
			service_file := service_path + "/" + field + make_file_suffix + ".go"
			fmt.Println("Create file :", service_file)
			fmt.Println("Create file :", service_file)
			fmt.Println("Create file :", service_file)
			file, err := os.OpenFile(service_file, os.O_CREATE|os.O_WRONLY, os.ModePerm)
			if err != nil {
				fmt.Println("open failed err:", err)
			} else {
				err = tmpl.Execute(file, data)
				fmt.Println("tmpl.Execute=>", err)
			}
			//break
		}
		c.HTML(http.StatusOK, "create success")
	}
	return nil
}
