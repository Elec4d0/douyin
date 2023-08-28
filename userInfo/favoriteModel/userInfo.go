package favoriteModel

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"userInfo/favoriteModel/api"
	"userInfo/favoriteModel/api/favoritemodelservice"
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

func QueryUserFavoriteCount(userID int64) (likeCount, totalFavorited int64, err error) {
	rpcReq := &api.FavoriteModelQueryUserFavoriteCountRequest{
		UserId: userID,
	}

	resp, err := favoriteModelRpcClient.QueryUserFavoriteCount(context.Background(), rpcReq)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	return resp.FavoriteCount, resp.TotalFavorited, nil
}

func BatchQueryUserFavoriteCount(userIDList []int64) (likeCountList, totalFavoritedList []int64, err error) {
	rpcReq := &api.FavoriteModelQueryUserFavoriteCountListRequest{
		UserIdList: userIDList,
	}

	resp, err := favoriteModelRpcClient.QueryUserFavoriteCountList(context.Background(), rpcReq)
	if err != nil || resp.TotalFavoritedList == nil || resp.FavoriteCountList == nil {
		log.Println(err)
		var emptyIntArr = make([]int64, len(userIDList))
		return emptyIntArr, emptyIntArr, err
	}

	return resp.FavoriteCountList, resp.TotalFavoritedList, nil
}
