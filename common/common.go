package common

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var Sep = string(os.PathSeparator)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// GetDateTime 获取当前日期时间
func GetDateTime() string {
	now := time.Now()
	dateString := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	return dateString
}

// InArrStr 判断指定数据是否在数组内
func InArrStr(arr []string, str string) bool {

	for _, val := range arr {
		if val == str {
			return true
		}
	}

	return false
}

/*
	PathExists
   判断文件或文件夹是否存在
   如果返回的错误为nil,说明文件或文件夹存在
   如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
   如果返回的错误为其它类型,则不确定是否在存在
*/
func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ListFiles(dirname string, level int) []string {

	var fileList = make([]string, 0)

	// level用来记录当前递归的层次
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + Sep + fi.Name()
		fileList = append(fileList, filename)
		if fi.IsDir() {
			//继续遍历fi这个目录
			fileList = append(fileList, ListFiles(filename, level+1)...)
		}
	}
	return fileList
}

// IsFile 判断是否是文件，不是文件就是目录
func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

// Md5File 计算一个文件的md5
func Md5File(filename string) (string, error) {
	f, err := os.Open(filename) //打开文件
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	defer f.Close()

	md5Handle := md5.New()         //创建 md5 句柄
	_, err = io.Copy(md5Handle, f) //将文件内容拷贝到 md5 句柄中
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	md := md5Handle.Sum(nil)        //计算 MD5 值，返回 []byte
	md5str := fmt.Sprintf("%x", md) //将 []byte 转为 string
	return md5str, nil
}
