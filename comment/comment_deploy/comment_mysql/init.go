package comment_mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 客户端运行前初始化
func Init() {
	//连接数据库
	dsn := "root:enid123456@tcp(127.0.0.1:3306)/commentMysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//自动创建数据库表
	err = db.AutoMigrate(&Comment{}, &CommentCount{})
	if err != nil {
		panic("failed to auto migrate database")
	}
	DB = db
}
