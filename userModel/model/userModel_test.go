package model

import (
	"fmt"
	"testing"
)

func TestFindUserByIDs(t *testing.T) {
	Init()
	var author_id []int64
	author_id = append(author_id, 1000009, 1000008, 1000)
	var userList []*User
	var err error
	userList, err = FindUserByIDs(author_id)
	for key, value := range userList {
		fmt.Println("key: ", key, "   ", "value: ", value)
	}
	fmt.Println("error: ", err)
}
