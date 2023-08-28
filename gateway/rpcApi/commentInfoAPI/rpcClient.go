package commentInfo

import (
	"context"
	"fmt"
	"gateway/rpcApi/commentInfoAPI/api"
	"gateway/rpcApi/commentInfoAPI/api/commentserver"
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

func GetCommentAllCount(videoIds []int64) (*api.DouyinCommentserverCommentallcountResponse, error) {
	rpcReq := &api.DouyinCommentserverCommentallcountRequest{
		VideoIds: videoIds,
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

func CommentAction(userId, videoId int64, actionType int32, commentText string, commentId int64) (*api.DouyinCommentActionResponse, error) {
	var resp *api.DouyinCommentActionResponse
	var err error
	if actionType == 1 {
		req := &api.DouyinCommentActionRequest{
			UserId:      userId,
			ActionType:  int32(actionType),
			VideoId:     int64(videoId),
			CommentText: &commentText,
		}
		log.Println(req)
		resp, err = commentInfoRpcClient.CommentAction(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
	} else if actionType == 2 {
		req := &api.DouyinCommentActionRequest{
			UserId:     userId,
			ActionType: int32(actionType),
			VideoId:    int64(videoId),
			CommentId:  &commentId,
		}
		log.Println(req)
		resp, err = commentInfoRpcClient.CommentAction(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		resp = nil
	}
	return resp, nil
}

func CommentList(userId, videoId int64) (*api.DouyinCommentListResponse, error) {
	req := &api.DouyinCommentListRequest{
		UserId:  userId,
		VideoId: videoId,
	}
	resp, err := commentInfoRpcClient.CommentList(context.Background(), req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return resp, nil
}
