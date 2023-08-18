package rpcClient

import (
	"fmt"
	"testing"
	userModelServices "userInfo/userModelAPI"
)

/*
func TestGetFullUserInfo(t *testing.T) {
	userModelServices.InitUserModelRpcClient()
	InitUserInfoRpcClient()
	fullUser, _ := GetFullUserInfo(1000008, 1000005)

	fmt.Println(fullUser)
}*/

func TestGetFullUserInfoList(t *testing.T) {
	userModelServices.InitUserModelRpcClient()
	InitUserInfoRpcClient()
	var search_id []int64
	search_id = append(search_id, 1000009, 1000008, 100008, 1000015)
	user, _ := GetFullUserInfoList(1000008, search_id)

	for i := 0; i < len(user); i++ {
		fmt.Println(user[i])
	}
}
