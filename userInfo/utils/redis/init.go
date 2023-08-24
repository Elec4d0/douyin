package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

var ctx = context.Background()

func Init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.41:6379",
		Password: "",
		DB:       5,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
	}
	fmt.Println("成功连接redis")
}

type User struct {
	Id              int64   `json:"id"`
	Name            string  `json:"name"`
	FollowCount     *int64  `json:"follow_count"`
	FollowerCount   *int64  `json:"follower_count"`
	Avatar          *string `json:"avatar"`
	BackgroundImage *string `json:"background_image"`
	Signature       *string `json:"signature"`
	TotalFavorited  *int64  `json:"total_favorited"`
	WorkCount       *int64  `json:"work_count"`
	FavoriteCount   *int64  `json:"favorite_count"`
}
