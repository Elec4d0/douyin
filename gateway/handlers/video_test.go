package handlers

import (
	"context"
	"fmt"
	"gateway/microService/feed/api"
	"testing"
)

func TestFeed(t *testing.T) {
	fmt.Println("开始测试——————TestFeed—————————————————")
	rpcClient := InitVideoRpcClient()

	var lastTime int64 = -1
	token := ""

	rpcReq := &api.DouyinFeedRequest{
		LatestTime: &lastTime,
		Token:      &token,
	}
	resp, err := rpcClient.GetFeed(context.Background(), rpcReq)
	fmt.Println(err)

	fmt.Println(resp)
	fmt.Println("结束测试——————TestFeed—————————————")
}
