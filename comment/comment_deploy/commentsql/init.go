package commentsql

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB
var rds redis.Conn

// 客户端运行前初始化
func MysqlInit() {
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

func RedisPollInit() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     5, //最大空闲数
		MaxActive:   0, //最大连接数
		Wait:        true,
		IdleTimeout: time.Duration(1) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			redis.DialDatabase(0)
			return c, err
		},
	}
}

func RedisInit() {
	rds = RedisPollInit().Get()
}

func RedisClose() {
	_ = rds.Close()
}
