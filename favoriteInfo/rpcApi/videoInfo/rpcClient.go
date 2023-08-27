package videoInfo

import (
	"context"
	"favoriteInfo/rpcApi/videoInfo/api"
	"favoriteInfo/rpcApi/videoInfo/api/videoinfoservice"
	"fmt"
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

func GetVideoInfoList(userID int64, videoIDList []int64) ([]*api.Video, error) {
	rpcReq := &api.VideoInfoGetVideoInfoListRequest{
		UserId:      userID,
		VideoIdList: videoIDList,
	}

	resp, err := videoInfoRpcClient.GetVideoInfoList(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp.VideoList, nil
}

func GetVideoInfo(userID int64, videoID int64) (*api.Video, error) {
	rpcReq := &api.VideoInfoGetVideoInfoRequest{
		UserId:  userID,
		VideoId: videoID,
	}

	resp, err := videoInfoRpcClient.GetVideoInfo(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp.Video, nil

}

func GetFeed(userID int64, nextTime int64) ([]*api.Video, error) {
	rpcReq := &api.VideoInfoGetFeedRequest{
		UserId:   userID,
		NextTime: nextTime,
	}

	resp, err := videoInfoRpcClient.GetFeed(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp.VideoList, nil

}

func GetAuthorVideoInfoList(authorID int64, userId int64) ([]*api.Video, error) {
	rpcReq := &api.VideoInfoGetAuthorVideoInfoListRequest{
		UserId:   userId,
		AuthorId: authorID,
	}

	resp, err := videoInfoRpcClient.GetAuthorVideoInfoList(context.Background(), rpcReq)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp.VideoList, nil
}
