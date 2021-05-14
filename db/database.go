package db

import (
	"filecloud/common"
	"filecloud/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DataBase struct {
	Db *gorm.DB
}

// Open 打开数据库
func (that *DataBase) Open(dbFile string) {
	that.Db, _ = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
}

func (that *DataBase) Find(dest interface{}) (tx *gorm.DB) {
	return that.Db.Find(dest)
}

// CreateFilesTable 创建文件表
func (that *DataBase) CreateFilesTable() {
	err := that.Db.AutoMigrate(&model.Files{})
	common.CheckErr(err)
}
