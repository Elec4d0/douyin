package videoInfo

import (
	"context"
	"fmt"
	"gateway/rpcApi/videoInfo/api"
	"gateway/rpcApi/videoInfo/api/videoinfoservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var videoInfoRpcClient videoinfoservice.Client

func InitVideoInfoRpcClient() videoinfoservice.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	videoInfoRpcClient, err = videoinfoservice.NewClient("videoInfo", client.WithResolver(r))

	if err != nil {
		log.Fatal("网关层Video 微服务初始化链接失败")
		log.Fatal(err)
		return nil
	}
	fmt.Println("Video 微服务：初始化链接User微服务成功")
	return videoInfoRpcClient
}

func GetVideoInfoList(videoIDList []int64) ([]*api.Video, error) {
	rpcReq := &api.VideoInfoGetVideoInfoListRequest{
		VideoIdList: videoIDList,
	}

	resp, err := videoInfoRpcClient.GetVideoInfoList(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp.VideoList, nil
}

func GetVideoInfo(videoID int64) (*api.Video, error) {
	rpcReq := &api.VideoInfoGetVideoInfoRequest{
		VideoId: videoID,
	}

	resp, err := videoInfoRpcClient.GetVideoInfo(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp.Video, nil

}

func GetFeed(userID int64, nextTime int64) (*api.VideoInfoGetFeedResponse, error) {
	rpcReq := &api.VideoInfoGetFeedRequest{
		UserId:   userID,
		NextTime: nextTime,
	}

	resp, err := videoInfoRpcClient.GetFeed(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp, nil

}

func GetAuthorVideoInfoList(authorID int64, userId int64) (*api.VideoInfoGetAuthorVideoInfoListResponse, error) {
	rpcReq := &api.VideoInfoGetAuthorVideoInfoListRequest{
		UserId:   userId,
		AuthorId: authorID,
	}

	resp, err := videoInfoRpcClient.GetAuthorVideoInfoList(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp, nil
}
