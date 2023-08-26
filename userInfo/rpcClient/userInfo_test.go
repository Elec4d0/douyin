package rpcClient

import (
	"fmt"
	"testing"
	userModelServices "userInfo/userModelAPI"
	videoModelServices "userInfo/videoModel"
)

/*
func TestGetFullUserInfo(t *testing.T) {
	userModelServices.InitUserModelRpcClient()
	videoModelServices.InitVideoModelRpcClient()
	InitUserInfoRpcClient()
	fullUser, _ := GetFullUserInfo(1000008, 1000009)

	fmt.Println(fullUser)
}*/

func TestGetFullUserInfoList(t *testing.T) {
	userModelServices.InitUserModelRpcClient()
	videoModelServices.InitVideoModelRpcClient()
	InitUserInfoRpcClient()
	var search_id []int64
	search_id = append(search_id, 1000009, 1000018)
	user, _ := GetFullUserInfoList(1000008, search_id)

	for i := 0; i < len(user); i++ {
		fmt.Println(user[i])
	}
}
