package upload

import (
	"filecloud/common"
	"filecloud/conf"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/golang/glog"
	"path/filepath"
)

type AliOSSFileCloud struct {
	client *oss.Client
	bucket *oss.Bucket
}

// Connect 连接至服务器
func (that *AliOSSFileCloud) Connect() {

	aliOSS := conf.Conf.AliOSS

	client, err := oss.New(aliOSS.Endpoint, aliOSS.AccessKeyID, aliOSS.AccessKeySecret)
	common.CheckErr(err)
	that.client = client

	if !that.isBucketExist(aliOSS.BucketName) {
		glog.Info("无法找到指定的bucketName[" + aliOSS.BucketName + "]")
	}
	bucket, err := client.Bucket(aliOSS.BucketName)
	common.CheckErr(err)
	that.bucket = bucket
	glog.Info("AliOSS Connect Success...")
}

func (that *AliOSSFileCloud) isBucketExist(bucketName string) bool {
	isExist, err := that.client.IsBucketExist(bucketName)
	common.CheckErr(err)
	return isExist
}
func (that *AliOSSFileCloud) FileExist(fileName string) bool {
	isExist, err := that.bucket.IsObjectExist(fileName)
	common.CheckErr(err)
	return isExist
}

func (that *AliOSSFileCloud) IsFolderExist(folderName string) bool {

	return false
}

func (that *AliOSSFileCloud) DataBaseExist() bool {
	return that.FileExist(conf.Conf.Upload.DbName)
}

func (that *AliOSSFileCloud) UploadFile(fileName string, localFileName string) error {
	return that.bucket.PutObjectFromFile(filepath.ToSlash(fileName), localFileName)
}

func (that *AliOSSFileCloud) DownLoadFile(fileName string, localFileName string) error {
	return that.bucket.GetObjectToFile(fileName, localFileName)
}
