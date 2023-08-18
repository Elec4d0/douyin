package rpcClient

import (
	"fmt"
	"testing"
	"userModel/model"
)

/*
func TestCreateBaseUser(t *testing.T) {
	InitUserModelRpcClient()

	id, _ := CreateBaseUser("hello_user", "123456")

	fmt.Println(id)
}

func TestFindBaseUserByName(t *testing.T) {
	InitUserModelRpcClient()
	user, _ := FindBaseUserByName("测试注册")

	fmt.Println(user.Id)
}

func TestFindBaseUserById(t *testing.T) {
	InitUserModelRpcClient()
	user, _ := FindBaseUserById(1000008)

	fmt.Println(user)
}*/

func TestFindBaseUserList(t *testing.T) {
	model.Init()
	InitUserModelRpcClient()
	var author_id []int64
	//for i := 0; i < 10; i++ {
	//	author_id = append(author_id, 1000000+int64(i))
	//}
	author_id = append(author_id, 1000009, 1000008, 100008, 1000015)
	author_id = append(author_id, 1000009)
	user, _ := FindBaseUserList(author_id)

	for _, value := range user {
		fmt.Println("ID: ", value.Id, "  user: ", value)
	}
}
