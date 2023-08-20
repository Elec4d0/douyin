package userInfo

import (
	"context"
	"fmt"
	api "gateway/rpcApi/userInfoAPI/api"
	userInfo "gateway/rpcApi/userInfoAPI/api"
	"gateway/rpcApi/userInfoAPI/api/userinfoservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
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

func GetFullUserInfo(user_id int64, search_id int64) (*userInfo.DouyinUserGetFullUserInfoResponse, error) {
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
	return resp, nil
}

func GetFullUserInfoList(user_id int64, search_id []int64) (*userInfo.DouyinUserGetFullUserInfoListResponse, error) {
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

	return resp, nil
}
