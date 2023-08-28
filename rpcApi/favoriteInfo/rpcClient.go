package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"rpcApi/favoriteInfo/api"
	"rpcApi/favoriteInfo/api/favoriteinfoservice"
)

var favoriteInfoRpcClient favoriteinfoservice.Client

func InitFavoriteModelRpcClient() favoriteinfoservice.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Println(err)
	}
	favoriteInfoRpcClient, err = favoriteinfoservice.NewClient("favoriteInfo", client.WithResolver(r))

	if err != nil {
		log.Println("favoriteInfo微服务rpcClient初始化链接失败")
		log.Println(err)
		return nil
	}
	fmt.Println("favoriteInfo微服务rpcClient初始化链接成功")
	return favoriteInfoRpcClient
}

func QueryFavoriteList(userID, searchID int64) ([]*api.FavoriteVideo, error) {
	rpcReq := &api.FavoriteInfoQueryFavoriteListRequest{UserId: userID, SearchId: searchID}

	resp, err := favoriteInfoRpcClient.QueryFavoriteList(context.Background(), rpcReq)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp.VideoList, err
}

func FavoriteAction(userID, videoID int64, actionType int32) (*api.FavoriteInfoFavoriteActionResponse, error) {
	rpcReq := &api.FavoriteInfoFavoriteActionRequest{
		UserId:     userID,
		VideoId:    videoID,
		ActionType: actionType,
	}

	resp, err := favoriteInfoRpcClient.FavoriteAction(context.Background(), rpcReq)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}
