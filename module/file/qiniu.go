package file

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"mime/multipart"
	"os"
	"strings"
	"time"
	"github.com/foxiswho/echo-go/module/cache"
	"github.com/foxiswho/echo-go/conf"
	"github.com/foxiswho/echo-go/util"
)

//七牛云存储
type QiNiu struct {
	Config *conf.Upload `json:"-"`
}

// 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

//初始化
func NewQiNiu() *QiNiu {
	return new(QiNiu)
}

//获取配置
func (t *QiNiu) setConfig() (bool, error) {
	t.Config = &conf.Conf.Upload
	return true, nil
}

//七牛配置读取
func (t *QiNiu) SetQiNiuConfig() (bool, error) {
	ok, err := t.setConfig()
	if err != nil {
		fmt.Println("setConfig err:", err)
		return false, err
	}
	fmt.Println("setConfig", ok)
	if t.Config != nil {
		return false, util.NewError("配置文件没有读取")
	}
	return true, nil
}

//设置 token 缓存
func (t *QiNiu) setToken() string {
	// 初始化AK，SK
	putPolicy := storage.PutPolicy{
		Scope:   t.Config.RootPath, // 设置上传到的空间 bucket
		Expires: 3601,               //设置Token过期时间
	}
	//access_key secret_key
	mac := qbox.NewMac(t.Config.RootPath, t.Config.RootPath)
	// 生成一个上传token
	token := putPolicy.UploadToken(mac)
	//缓存
	err := cache.Client().Set("qiniu_token", token, 3600*time.Second)
	if err != nil {
		fmt.Println("设置缓存错误", err)
	}
	return token
}

//获取
func (t *QiNiu) GetToken() string {
	str := ""
	err := cache.Client().Get("qiniu_token", &str)
	if err != nil {
		fmt.Println("获取缓存错误", err)
	}
	if len(str) > 0 {
		return str
	}
	return t.setToken()
}

//上传
func (t *QiNiu) Upload(file multipart.File, UploadFile *UploadFile) (interface{}, error) {
	//七牛配置填充
	_, err := t.SetQiNiuConfig()
	if err != nil {
		fmt.Println("getConfig err:", err)
		return nil, util.NewError("七牛配置错误:" + err.Error())
	}
	//token
	token := t.GetToken()
	if len(token) < 1 {
		fmt.Println("token", token)
		return nil, util.NewError("七牛token不能为空")
	}
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}

	//当前的目录
	dir, _ := os.Getwd()
	fmt.Println("当前的目录", dir)

	// 设置上传文件的路径
	key := UploadFile.Path + UploadFile.Name
	filePath := dir + UploadFile.GetLocalTmpPath() + UploadFile.Name
	key = strings.Replace(key, "/", "", 1)
	fmt.Println("本地文件绝对路径", filePath)
	fmt.Println("去除第一个字符/后，访问路径", key)
	// 调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	err = formUploader.PutFile(context.Background(), &ret, token, key, filePath, &putExtra)
	// 打印出错信息
	if err != nil {
		fmt.Println("io.Put failed:", err)
		return nil, err
	}
	// 打印返回的信息
	fmt.Println("七牛成功后返回信息", ret)
	return ret, nil
}
