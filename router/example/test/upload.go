package test

import (
	"github.com/foxiswho/echo-go/router/base"
	"os"
	"io"
	"net/http"
	"fmt"
	"github.com/foxiswho/echo-go/conf"
	"github.com/foxiswho/echo-go/module/file"
	"github.com/foxiswho/echo-go/util/conv"
)

type Upload struct {
}

func NewUpload() *Upload {
	return new(Upload)
}
func (x *Upload) UploadIndex(c *base.BaseContext) error {
	//上传令牌 初始化
	maps := make(map[string]interface{})
	maps["type_id"] = 1
	maps["id"] = 1
	maps["aid"] = 1
	cry, err := file.TokeMake(maps)
	if err != nil {
		fmt.Println("令牌加密错误：" + err.Error())
		c.Error(err)
	}
	c.Set("tmpl", "example/test/upload")
	c.Set("data", map[string]interface{}{
		"title": "上传",
		"upload_token": cry,
	})
	return nil
}

func UploadPostIndex(c *base.BaseContext) error {

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	root_path := "." + conf.Conf.Upload.RootPath
	err = os.MkdirAll(root_path, os.ModePerm)
	fmt.Print("Create Directory=========", root_path)
	if err != nil {
		fmt.Printf("Create Directory ERROR %s", err)
	} else {
		fmt.Print("Create Directory OK! ", root_path)
	}
	// Destination
	dst, err := os.Create(root_path + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully </p>", root_path+file.Filename))
}

func UploadMorePostIndex(c *base.BaseContext) error {
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]
	root_path := "." + conf.Conf.Upload.RootPath
	err = os.MkdirAll(root_path, os.ModePerm)
	fmt.Print("Create Directory=========", root_path)
	if err != nil {
		fmt.Printf("Create Directory ERROR %s", err)
	} else {
		fmt.Print("Create Directory OK! ", root_path)
	}
	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(root_path + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files .</p>", len(files)))
}

//上传图片 支持 markdown编辑器上传图片
// @router /upload/image [post]
func UploadDbHandler(c *base.BaseContext) error {
	fmt.Println("pppppppppp")
	fmt.Println("pppppppppp")
	fmt.Println("pppppppppp")
	fmt.Println("pppppppppp")
	fmt.Println("pppppppppp")
	pp,err:=c.FormParams()
	fmt.Println("pppppppppp",pp,err)
	//声明
	var maps map[string]interface{}
	t := c.FormValue("t")
	fmt.Println("token=>",c.FormValue("token"))
	token := c.FormValue("token")
	fmt.Println("token=>",token)
	//token 验证
	if len(token) > 0 {
		//解密
		maps, err = file.TokenDeCode(token)
		if err != nil {
			fmt.Println("令牌：" + token)
			fmt.Println("令牌解密失败：" + err.Error())
			token = ""
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "令牌解密失败：" + err.Error(),
			})
			return nil
		} else {
			fmt.Println("令牌解密", maps)
		}
	}
	//判断是否是 markdown编辑器 输出相应的错误
	if token == "" {
		if t == "markdown" {
			//md := &editor.EditorMd{}
			//md.Message = "令牌错误"
			//md.Success = 0
			//c.JSON(http.StatusBadRequest, map[string]interface{}{
			//	"json": md,
			//})
		} else {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "令牌错误：",
			})
		}
		return nil
	}
	//上传文件file表单元素名称
	file_name := "file"
	if t == "markdown" {
		file_name = "editormd-image-file"
	}
	//上传
	f, err := file.Upload(file_name, c.Request(), maps)
	//如果是markdown编辑器返回
	if t == "markdown" {
		//md := &editor.EditorMd{}
		//if err != nil {
		//	md.Message = err.Error()
		//	md.Success = 0
		//} else {
		//	md.Message = "上传成功"
		//	md.Url = f.Http
		//	md.Success = 1
		//}
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "",
			//"json":    md,
		})
	} else {
		//其他返回
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		} else {
			m, err := conv.ObjToMap(f)
			if err != nil {
				c.JSON(http.StatusOK, map[string]interface{}{
					"message": "操作成功",
				})
			} else {
				c.JSON(http.StatusOK, map[string]interface{}{
					"message": "操作成功",
					"json":    m,
				})
			}
		}

	}
	return nil
}
