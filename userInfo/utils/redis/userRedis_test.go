package redis

import (
	"fmt"
	"testing"
	"userInfo/rpcClient"
	"userInfo/services/protos/kitex_gen/api"
	userModelServices "userInfo/userModelAPI"
	videoModelServices "userInfo/videoModel"
)

/*
func TestInsertUserRedis(t *testing.T) {
	userModelServices.InitUserModelRpcClient()
	videoModelServices.InitVideoModelRpcClient()
	rpcClient.InitUserInfoRpcClient()
	Init()
	var str string
	str = ""
	var a int64
	a = 0
	data := &api.FullUser{
		Id:              100016,
		Name:            "我是0",
		FollowCount:     &a,
		FollowerCount:   &a,
		Avatar:          &str,
		BackgroundImage: &str,
		Signature:       &str,
		TotalFavorited:  &a,
		WorkCount:       &a,
		FavoriteCount:   &a,
		IsFollow:        false,
	}
	_ = InsertUserRedis(data)

	user, _ := FindUserRedis(100016)
	fmt.Println(user)
}*/

func TestInsertUserListRedis(t *testing.T) {
	userModelServices.InitUserModelRpcClient()
	videoModelServices.InitVideoModelRpcClient()
	rpcClient.InitUserInfoRpcClient()
	Init()
	var str string
	str = ""
	var a int64
	a = 0
	data := &api.FullUser{
		Id:              100016,
		Name:            "我是0",
		FollowCount:     &a,
		FollowerCount:   &a,
		Avatar:          &str,
		BackgroundImage: &str,
		Signature:       &str,
		TotalFavorited:  &a,
		WorkCount:       &a,
		FavoriteCount:   &a,
		IsFollow:        false,
	}
	var userList []*api.FullUser
	userList = append(userList, data)
	userList = append(userList, data)
	_ = InsertUserListRedis(userList)

	var id []int64
	id = append(id, 100016, 100015)
	user, _ := FindUserListRedis(id)
	for _, value := range user {
		fmt.Println(value)
	}
}
