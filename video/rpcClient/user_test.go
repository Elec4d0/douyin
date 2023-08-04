package rpcClient

import (
	"fmt"
	"testing"
)

func TestGetUserInFo(t *testing.T) {
	InitUserRpcClient()
	fmt.Println("开始测试——————TestQuerySingleVideo——————————————————")
	var uId uint64 = 123123
	tmp, err := GetUserInFo(uId)
	if err != nil {
		fmt.Println("查询失败")
	} else {
		fmt.Println("查询成功，video对象结果：")

	}
	fmt.Println(tmp)
	fmt.Println("结束测试——————TestQuerySingleVideo——————————————————")
}
