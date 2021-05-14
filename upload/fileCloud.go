package upload

// FileCloud 所有云端通用的接口
type FileCloud interface {
	// Connect 连接至服务器
	Connect()
	// DataBaseExist 判断文件传输的数据库是否存在
	DataBaseExist() bool
	// IsFolderExist 判断文件夹是否存在
	IsFolderExist(folderName string) bool
	// FileExist 判断指定文件是否存在
	FileExist(fileName string) bool
	// UploadFile 上传文件至云端
	UploadFile(fileName string, localFileName string) error
	// DownLoadFile 下载云端文件到本地
	DownLoadFile(fileName string, localFileName string) error
}
