package conf

import (
	"encoding/json"
	"filecloud/common"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
)

// 公开的配置项
var (
	Conf           = &Config{}
	confPath       string
	uploadTypeList = []string{"AliOSS", "FTP"}
)

// Config .
type Config struct {
	Upload Upload
	AliOSS AliOSS
}

// AliOSS 阿里云OSS的配置
type AliOSS struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
}

// Upload 上传的相关配置
type Upload struct {
	Type     string
	DbName   string
	BasePath string
}

func init() {
	flag.StringVar(&confPath, "c", "./conf/conf.toml", "-c path")

	_, err := toml.DecodeFile(confPath, &Conf)

	common.CheckErr(err)

	if !common.InArrStr(uploadTypeList, Conf.Upload.Type) {
		jsonStr, _ := json.Marshal(uploadTypeList)
		panic(fmt.Sprintf("暂不支持您提供的文件上传方式,暂时仅支持：%s\n", jsonStr))
	}

}
