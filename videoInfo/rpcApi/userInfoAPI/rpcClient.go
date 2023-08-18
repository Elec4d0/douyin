package rpcClient

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"videoInfo/rpcApi/userInfoAPI/api"
	"videoInfo/rpcApi/userInfoAPI/api/userinfoservice"
)

var userInfoRpcClient userinfoservice.Client

func InitUserInfoRpcClient() userinfoservice.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	userInfoRpcClient, err = userinfoservice.NewClient("userInfoService", client.WithResolver(r))

	if err != nil {
		log.Println("网关层userInfo 微服务初始化链接失败")
		log.Fatal(err)
		return nil
	}
	fmt.Println("userInfo 微服务：初始化链接成功")
	return userInfoRpcClient
}

func GetFullUserInfo(user_id int64, search_id int64) (*api.FullUser, error) {
	rpcReq := &api.DouyinUserGetFullUserInfoRequest{
		UserId:   user_id,
		SearchId: search_id,
	}
	fmt.Println(rpcReq)
	resp, err := userInfoRpcClient.GetFullUserInfo(context.Background(), rpcReq)

	if err != nil {
		log.Println(resp)
		log.Println(err)
		return nil, err
	}
	return resp.FullUser, nil
}

func GetFullUserInfoList(user_id int64, search_id []int64) ([]*api.FullUser, error) {
	rpcReq := &api.DouyinUserGetFullUserInfoListRequest{
		UserId:   user_id,
		SearchId: search_id,
	}
	fmt.Println(rpcReq)
	resp, err := userInfoRpcClient.GetFullUserInfoList(context.Background(), rpcReq)

	if err != nil {
		log.Println(resp)
		log.Println(err)
		return nil, err
	}
	return resp.FullUser, nil
}
