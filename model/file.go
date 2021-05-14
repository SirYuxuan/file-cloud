package model

type Files struct {
	Id         int `gorm:"column:id;not null;type:integer primary key autoincrement;"`
	FileName   string
	Md5        string
	CreateTime string
}

// Insert 插入一条文件记录到数据库中
func (that *Files) Insert() {

}
