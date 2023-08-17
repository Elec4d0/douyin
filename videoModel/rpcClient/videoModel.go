package rpcClient

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"videoModel/core/kitex_gen/api"
	"videoModel/core/kitex_gen/api/videomodelservice"
)

var videoModelRpcClient videomodelservice.Client

func InitVideoModelRpcClient() videomodelservice.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	videoModelRpcClient, err = videomodelservice.NewClient("videoModel", client.WithResolver(r))

	if err != nil {
		log.Fatal("网关层Video 微服务初始化链接失败")
		log.Fatal(err)
		return nil
	}
	fmt.Println("Video 微服务：初始化链接User微服务成功")
	return videoModelRpcClient
}

func CreateVideo(AuthorId int64, PlayUrl string, CoverUrl string, Title string) error {
	rpcReq := &api.VideoModelCreateVideoRequest{
		AuthorId: AuthorId,
		PlayUrl:  PlayUrl,
		CoverUrl: CoverUrl,
		Title:    Title,
	}
	fmt.Println(rpcReq)
	resp, err := videoModelRpcClient.CreateVideo(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		log.Fatal(resp)
		return err
	}
	return nil
}

func QueryAuthorWorkCount(AuthorId int64) (int64, error) {
	rpcReq := &api.VideoModelQueryAuthorWorkCountRequest{
		AuthorId: AuthorId,
	}

	resp, err := videoModelRpcClient.QueryAuthorWorkCount(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return int64(resp.WorkCount), nil
}

func QueryAuthorVideoList(AuthorId int64) ([]*api.VideoBaseInfo, error) {
	rpcReq := &api.VideoModelQueryAuthorVideoListRequest{
		AuthorId: AuthorId,
	}

	resp, err := videoModelRpcClient.QueryAuthorVideoList(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp.VideoBaseInfo, nil
}

func QueryVideoList(videoIDs []int64) ([]*api.VideoBaseInfo, error) {
	rpcReq := &api.VideoModelQueryVideoListRequest{
		VideoIdList: videoIDs,
	}

	resp, err := videoModelRpcClient.QueryVideoList(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp.VideoBaseInfoList, nil
}

func QueryVideo(videoID int64) (*api.VideoBaseInfo, error) {
	rpcReq := &api.VideoModelQueryVideoRequest{
		VideoId: videoID,
	}

	resp, err := videoModelRpcClient.QueryVideo(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp.VideoBaseInfo, nil
}

func QueryVideoFeed(nextTime int64) ([]*api.VideoBaseInfo, error) {
	rpcReq := &api.VideoModelQueryVideoFeedRequest{
		NextTime: nextTime,
	}

	resp, err := videoModelRpcClient.QueryVideoFeed(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp.VideoBaseInfoList, nil
}
