package main

import (
	"filecloud/common"
	"filecloud/conf"
	"filecloud/db"
	"filecloud/model"
	"filecloud/upload"
	"github.com/golang/glog"
	"gorm.io/gorm"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Server 核心主服务
type Server struct {
	FileCloud  upload.FileCloud
	uploadConf conf.Upload
	Database   db.DataBase
}

// Build 完成所有初始化工作并监听文件变动
func (that *Server) Build() {

	that.uploadConf = conf.Conf.Upload

	if strings.ToUpper(that.uploadConf.Type) == "ALIOSS" {
		that.FileCloud = &upload.AliOSSFileCloud{}
	} else if strings.ToUpper(that.uploadConf.Type) == "FTP" {
		that.FileCloud = &upload.FtpFileCloud{}
	}
	that.FileCloud.Connect()

	that.InitDataBase()

	that.SynLocalFile()

}

// InitDataBase 初始化数据库
func (that *Server) InitDataBase() {

	dbName := that.uploadConf.BasePath + that.uploadConf.DbName

	// 核心数据库不存在 准备创建
	if !that.FileCloud.DataBaseExist() {
		// 判断本地是否存在，如果本地存在则上传到云端
		exist, err := common.PathExists(dbName)
		if (err != nil && os.IsNotExist(err)) || !exist {
			// 本地也不存在
			that.Database.Open(dbName)
			that.Database.CreateFilesTable()
			glog.Info("创建数据库文件,", dbName)

		}
		// 本地存在 上传云端
		err = that.FileCloud.UploadFile(path.Base(filepath.ToSlash(dbName)), dbName)
		common.CheckErr(err)
		glog.Info("数据库文件初始化完毕,本地上传至云端")
	} else {
		exist, err := common.PathExists(dbName)
		if err == nil && exist {
			// 本地存在 删除本地数据库
			err = os.Remove(dbName)
			common.CheckErr(err)
		}
		err = that.FileCloud.DownLoadFile(that.uploadConf.DbName, dbName)
		common.CheckErr(err)
		that.Database.Open(dbName)
		glog.Info("数据库文件初始化完毕,云端同步至本地")
	}
}

// SynLocalFile 同步本地文件到云端
func (that *Server) SynLocalFile() {
	// 删除全部记录 开始同步
	that.Database.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Files{})
	fileList := common.ListFiles(conf.Conf.Upload.BasePath, 0)

	for _, val := range fileList {

		if common.IsFile(val) {
			fileName := strings.Replace(val, that.uploadConf.BasePath, "", -1)
			if !that.FileCloud.FileExist(filepath.ToSlash(fileName[1:])) {
				err := that.FileCloud.UploadFile(fileName[1:], val)
				common.CheckErr(err)

				md5, md5Err := common.Md5File(val)
				common.CheckErr(md5Err)

				file := &model.Files{
					FileName:   fileName[1:],
					CreateTime: common.GetDateTime(),
					Md5:        md5,
				}

				that.Database.Db.Create(file)

				glog.Info("文件[" + fileName[1:] + "],MD5[" + md5 + "]已上传")
			}

		}

	}

}

// SynCloudFile 同步云端文件到本地
func (that *Server) SynCloudFile() {

}
