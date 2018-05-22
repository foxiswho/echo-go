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
)

func ServiceMakeHandler(c *base.BaseContext) error {

	sql := "show tables"
	result, err := db.DB().Engine.QueryString(sql)
	fmt.Println("err", err)
	fmt.Println("result", result)
	if err != nil {
		fmt.Println("err", err)
	} else {
		template_file := "./template/design/make/service.go.tpl"
		service_path := "./service"
		field := "Tables_in_" + conf.Conf.DB.Name
		for i, val := range result {
			fmt.Println("result index=>", i)
			fmt.Println("result val=>", val)
			fmt.Println("result val=>", val[field])
			tmpl, err := template.ParseFiles(template_file)
			fmt.Println("template err", err)
			fmt.Println("template err", tmpl)
			//val[field] = "admin_menu"
			data := make(map[string]interface{})
			data["tables"] = val[field]
			//data["tables_first"] = gonicCasedName(val[field])
			data["tables_camel_case"] = str.BigCamelCase(val[field])        //驼峰命名法 大写
			data["tables_big_camel_case"] = data["tables_camel_case"]        //驼峰命名法 大写
			data["tables_little_camel_case"] =str.LittleCamelCase(val[field]) //驼峰命名法 小写
			fmt.Println("data=>", data)
			//
			//
			//
			err = os.MkdirAll(service_path, os.ModePerm)
			if err != nil {
				fmt.Println("%s", err)
			} else {
				fmt.Println("Create Directory OK! ", service_path)
			}
			service_file := service_path + "/" + val[field] + "_auto_make.go"
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