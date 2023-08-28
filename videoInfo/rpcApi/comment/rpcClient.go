package rpcClient

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"videoInfo/rpcApi/comment/api"
	"videoInfo/rpcApi/comment/api/commentserver"
)

var commentInfoRpcClient commentserver.Client

func InitCommentInfoRpcClient() commentserver.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	commentInfoRpcClient, err = commentserver.NewClient("CommentServer", client.WithResolver(r))
	if err != nil {
		log.Println("comment微服务初始化链接失败")
		log.Fatal(err)
		return nil
	}
	fmt.Println("comment微服务：初始化链接成功")
	return commentInfoRpcClient
}

func GetCommentCount(video_id int64) (int64, error) {
	rpcReq := &api.DouyinCommentserverCommentcountRequest{
		VideoId: video_id,
	}

	resp, err := commentInfoRpcClient.CommentCount(context.Background(), rpcReq)

	if err != nil {
		log.Println(err)
		return 0, err
	}
	return resp.CommentCount, nil
}

func GetCommentAllCount(video_ids []int64) ([]int64, error) {
	rpcReq := &api.DouyinCommentserverCommentallcountRequest{
		VideoIds: video_ids,
	}

	resp, err := commentInfoRpcClient.CommentAllCount(context.Background(), rpcReq)
	if err != nil || resp.CommentCounts == nil {
		log.Println(err)
		var emptyIntArr = make([]int64, len(video_ids))
		return emptyIntArr, err
	}
	return resp.CommentCounts, nil
}
