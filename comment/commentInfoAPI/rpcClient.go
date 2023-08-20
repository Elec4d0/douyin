package commentInfo

import (
	"comment/server/protos/kitex_gen/api"
	"comment/server/protos/kitex_gen/api/commentserver"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var commentInfoRpcClient commentserver.Client

func InitCommentInfoRpcClient() commentserver.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	commentInfoRpcClient, err = commentserver.NewClient("CommentServer", client.WithResolver(r))
	if err != nil {
		log.Println("网关层comment微服务初始化链接失败")
		log.Fatal(err)
		return nil
	}
	fmt.Println("comment微服务：初始化链接成功")
	return commentInfoRpcClient
}

func GetCommentCount(video_id int64) (*api.DouyinCommentserverCommentcountResponse, error) {
	rpcReq := &api.DouyinCommentserverCommentcountRequest{
		VideoId: video_id,
	}
	fmt.Println(rpcReq)
	resp, err := commentInfoRpcClient.CommentCount(context.Background(), rpcReq)
	if err != nil {
		log.Println(resp)
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

func GetCommentAllCount(video_ids []int64) (*api.DouyinCommentserverCommentallcountResponse, error) {
	rpcReq := &api.DouyinCommentserverCommentallcountRequest{
		VideoIds: video_ids,
	}
	fmt.Println(rpcReq)
	resp, err := commentInfoRpcClient.CommentAllCount(context.Background(), rpcReq)
	if err != nil {
		log.Println(resp)
		log.Println(err)
		return nil, err
	}
	return resp, nil
}
