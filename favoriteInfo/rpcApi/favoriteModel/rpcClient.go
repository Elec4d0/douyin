package favoriteModel

import (
	"context"
	"favoriteInfo/rpcApi/favoriteModel/api"
	"favoriteInfo/rpcApi/favoriteModel/api/favoritemodelservice"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
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

func QueryFavoriteList(userID int64) ([]int64, error) {
	rpcReq := &api.FavoriteModelQueryFavoriteListRequest{UserId: userID}

	resp, err := favoriteModelRpcClient.QueryFavoriteList(context.Background(), rpcReq)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp.VideoIdList, err
}

func FavoriteAction(userID, videoID, authorID int64, actionType int32) error {
	rpcReq := &api.FavoriteModelFavoriteActionRequest{
		UserId:     userID,
		VideoId:    videoID,
		AuthorId:   authorID,
		ActionType: actionType,
	}

	_, err := favoriteModelRpcClient.FavoriteAction(context.Background(), rpcReq)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
