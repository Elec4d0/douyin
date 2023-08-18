package rpcClient

import (
	"fmt"
	"testing"
)

func TestCreateVideo(t *testing.T) {
	InitVideoModelRpcClient()

	CreateVideo(1, "http", "http", "测试标题")

	count, _ := QueryAuthorWorkCount(1)
	fmt.Println(count)
}

func TestQueryAuthorVideoList(t *testing.T) {
	res, _ := QueryAuthorVideoList(1)
	fmt.Println(res)
}

func TestQueryVideoList(t *testing.T) {

	var list = []int64{1, 38, 3, 4, 5}

	res, _ := QueryVideoList(list)
	fmt.Println(res)
}
