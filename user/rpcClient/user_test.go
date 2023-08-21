package rpcClient

import (
	"fmt"
	"testing"
)

func TestUserLogin(t *testing.T) {
	InitUserRpcClient()
	id, token, _ := UserLogin("lyy", "123456")
	fmt.Println("id: ", id, "   token:  ", token)
	id, token, _ = UserLogin("", "")
	fmt.Println("id: ", id, "   token:  ", token)
	id, token, _ = UserLogin("在山的那边", "123456")
	fmt.Println("id: ", id, "   token:  ", token)
	id, token, _ = UserLogin("在山的那边", "12345")
	fmt.Println("id: ", id, "   token:  ", token)
}

func TestUserRegister(t *testing.T) {
	InitUserRpcClient()
	id, token, _ := UserRegister("lyy", "123456")
	fmt.Println("id: ", id, "   token:  ", token)
	id, token, _ = UserRegister("南宁包司令官方", "123456")
	fmt.Println("id: ", id, "   token:  ", token)
	id, token, _ = UserRegister("", "")
	fmt.Println("id: ", id, "   token:  ", token)
}
