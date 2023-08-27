package favoriteModel

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"videoInfo/rpcApi/favoriteModel/api"
	"videoInfo/rpcApi/favoriteModel/api/favoritemodelservice"
)

var favoriteModelRpcClient favoritemodelservice.Client

func InitFavoriteModelRpcClient() favoritemodelservice.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Println(err)
	}
	favoriteModelRpcClient, err = favoritemodelservice.NewClient("favoriteModel", client.WithResolver(r))

	if err != nil {
		log.Println("favoriteModel微服务rpcClient初始化链接失败")
		log.Println(err)
		return nil
	}
	fmt.Println("favoriteModel微服务rpcClient初始化链接成功")
	return favoriteModelRpcClient
}

func QueryVideoFavoriteCount(videoID int64) (int64, error) {
	rpcReq := &api.FavoriteModelQueryVideoFavoriteCountRequest{
		VideoId: videoID,
	}

	resp, err := favoriteModelRpcClient.QueryVideoFavoriteCount(context.Background(), rpcReq)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return resp.VideoFavoriteCount, nil
}

func BatchQueryVideoFavoriteCount(videoIDList []int64) ([]int64, error) {
	rpcReq := &api.FavoriteModelQueryVideoFavoriteCountListRequest{
		VideoIdList: videoIDList,
	}

	resp, err := favoriteModelRpcClient.QueryVideoFavoriteCountList(context.Background(), rpcReq)

	if err != nil || resp.VideoFavoriteCountList == nil {
		log.Println(err)
		var emptyIntArr = make([]int64, len(videoIDList))
		return emptyIntArr, err
	}

	return resp.VideoFavoriteCountList, nil
}

func QueryIsUserFavorite(userID, videoID int64) (bool, error) {
	rpcReq := &api.FavoriteModelQueryIsUserFavoriteRequest{
		UserId:  userID,
		VideoId: videoID,
	}

	resp, err := favoriteModelRpcClient.QueryIsUserFavorite(context.Background(), rpcReq)

	if err != nil {
		log.Println(err)
		return false, err
	}
	return resp.IsUserFavorite, nil
}

func BatchQueryIsUserFavorite(userID int64, videoIDList []int64) ([]bool, error) {
	rpcReq := &api.FavoriteModelQueryIsUserFavoriteListRequest{
		UserId:      userID,
		VideoIdList: videoIDList,
	}
	resp, err := favoriteModelRpcClient.QueryIsUserFavoriteList(context.Background(), rpcReq)

	if err != nil || resp.IsUserFavoriteList == nil {
		log.Println(err)
		var emptyBoolArr = make([]bool, len(videoIDList))
		return emptyBoolArr, err
	}
	return resp.IsUserFavoriteList, nil
}
