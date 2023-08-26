package videoModel

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"videoInfo/rpcApi/videoModel/api"
	"videoInfo/rpcApi/videoModel/api/videomodelservice"
)

var videoModelRpcClient videomodelservice.Client

func InitVideoModelRpcClient() videomodelservice.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Println(err)
	}
	videoModelRpcClient, err = videomodelservice.NewClient("videoModel", client.WithResolver(r))

	if err != nil {
		log.Println("videoModel微服务rpcClient初始化链接失败")
		log.Println(err)
		return nil
	}
	fmt.Println("videoModel微服务rpcClient初始化链接成功")
	return videoModelRpcClient
}

func CreateVideo(AuthorId int64, PlayUrl string, CoverUrl string, Title string) error {
	rpcReq := &api.VideoModelCreateVideoRequest{
		AuthorId: AuthorId,
		PlayUrl:  PlayUrl,
		CoverUrl: CoverUrl,
		Title:    Title,
	}

	resp, err := videoModelRpcClient.CreateVideo(context.Background(), rpcReq)

	if err != nil {
		log.Println(err)
		log.Println(rpcReq)
		log.Println(resp)
		return err
	}
	return nil
}

func QueryAuthorWorkCount(AuthorId int64) (int32, error) {
	rpcReq := &api.VideoModelQueryAuthorWorkCountRequest{
		AuthorId: AuthorId,
	}

	resp, err := videoModelRpcClient.QueryAuthorWorkCount(context.Background(), rpcReq)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return resp.WorkCount, nil
}

func QueryAuthorVideoIDList(AuthorId int64) ([]int64, error) {
	rpcReq := &api.VideoModelQueryAuthorVideoIdListRequest{
		AuthorId: AuthorId,
	}

	resp, err := videoModelRpcClient.QueryAuthorVideoIDList(context.Background(), rpcReq)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp.VideoIdList, nil
}

func QueryVideoList(videoIDs []int64) ([]*api.VideoModel, error) {
	rpcReq := &api.VideoModelQueryVideoListRequest{
		VideoIdList: videoIDs,
	}

	resp, err := videoModelRpcClient.QueryVideoList(context.Background(), rpcReq)

	if err != nil {
		log.Println(err)
		return make([]*api.VideoModel, len(videoIDs)), err
	}

	return resp.VideoModelList, nil
}

func QueryVideo(videoID int64) (*api.VideoModel, error) {
	rpcReq := &api.VideoModelQueryVideoRequest{
		VideoId: videoID,
	}

	resp, err := videoModelRpcClient.QueryVideo(context.Background(), rpcReq)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp.VideoModel, nil
}

func QueryVideoFeed(nextTime int64, limit int64) (videoIDList []int64, createTimeList []int64, err error) {
	rpcReq := &api.VideoModelQueryVideoFeedRequest{
		NextTime: nextTime,
		Limit:    limit,
	}

	resp, err := videoModelRpcClient.QueryVideoFeed(context.Background(), rpcReq)

	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	return resp.VideoIdList, resp.CreateTimeList, nil
}

func QueryAuthorWorkCountList(authorIDList []int64) ([]int32, error) {
	rpcReq := &api.VideoModelQueryAuthorWorkCountListRequest{
		AuthorIdList: authorIDList,
	}

	resp, err := videoModelRpcClient.QueryAuthorWorkCountList(context.Background(), rpcReq)
	if err != nil {
		log.Println(err)
		return make([]int32, len(authorIDList)), err
	}

	return resp.WorkCountList, nil
}
