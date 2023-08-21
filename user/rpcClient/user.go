package rpcClient

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"user/services/protos/kitex_gen/api"
	"user/services/protos/kitex_gen/api/userservice"
)

var userRpcClient userservice.Client

func InitUserRpcClient() userservice.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	userRpcClient, err = userservice.NewClient("userService", client.WithResolver(r))

	if err != nil {
		log.Println("网关层user 微服务初始化链接失败")
		log.Fatal(err)
		return nil
	}
	fmt.Println("user 微服务：初始化链接成功")
	return userRpcClient
}

func UserLogin(name string, password string) (int64, string, error) {
	rpcReq := &api.DouyinUserLoginRequest{
		Username: name,
		Password: password,
	}
	fmt.Println(rpcReq)
	resp, err := userRpcClient.UserLogin(context.Background(), rpcReq)

	if err != nil {
		log.Println(resp)
		log.Println(err)
		return -1, "", err
	}
	return resp.UserId, resp.Token, nil
}

func UserRegister(name string, password string) (int64, string, error) {
	rpcReq := &api.DouyinUserRegisterRequest{
		Username: name,
		Password: password,
	}
	fmt.Println(rpcReq)
	resp, err := userRpcClient.UserRegister(context.Background(), rpcReq)

	if err != nil {
		log.Println(resp)
		log.Println(err)
		return -1, "", err
	}
	return resp.UserId, resp.Token, nil
}
