package upload

type FtpFileCloud struct {
}

// Connect 连接至服务器
func (that *FtpFileCloud) Connect() {

}

func (that *FtpFileCloud) isBucketExist(bucketName string) bool {

	return false
}
func (that *FtpFileCloud) FileExist(fileName string) bool {

	return false
}

func (that *FtpFileCloud) IsFolderExist(folderName string) bool {

	return false
}

func (that *FtpFileCloud) DataBaseExist() bool {
	return that.FileExist("database.db")
}

func (that *FtpFileCloud) UploadFile(fileName string, localFileName string) error {
	return nil
}

func (that *FtpFileCloud) DownLoadFile(fileName string, localFileName string) error {
	return nil
}
