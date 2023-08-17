package rpcClient

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"userModel/services/protos/kitex_gen/api"
	"userModel/services/protos/kitex_gen/api/usermodelservice"
)

var userModelRpcClient usermodelservice.Client

func InitUserModelRpcClient() usermodelservice.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	userModelRpcClient, err = usermodelservice.NewClient("UserModelService", client.WithResolver(r))

	if err != nil {
		log.Fatal("网关层userModel 微服务初始化链接失败")
		log.Fatal(err)
		return nil
	}
	fmt.Println("userModel 微服务：初始化链接成功")
	return userModelRpcClient
}

func CreateBaseUser(username string, password string) (int64, error) {
	rpcReq := &api.DouyinUserCreateBaseUserRequest{
		Username: username,
		Password: password,
	}
	fmt.Println(rpcReq)
	resp, err := userModelRpcClient.CreateBaseUser(context.Background(), rpcReq)

	if err != nil {
		log.Println(resp)
		log.Fatal(err)
		return 0, err
	}
	return resp.UserId, nil
}

func FindBaseUserByName(username string) (*api.BaseUser, error) {
	rpcReq := &api.DouyinUserFindBaseUserByNameRequest{
		Username: username,
	}
	fmt.Println(rpcReq)
	resp, err := userModelRpcClient.FindBaseUserByName(context.Background(), rpcReq)

	if err != nil {
		log.Println(resp)
		log.Fatal(err)
		return nil, err
	}
	return resp.BaseUser, nil
}

func FindBaseUserById(user_id int64) (*api.BaseUser, error) {
	rpcReq := &api.DouyinUserFindBaseUserByIdRequest{
		UserId: user_id,
	}
	fmt.Println(rpcReq)
	resp, err := userModelRpcClient.FindBaseUserById(context.Background(), rpcReq)

	if err != nil {
		log.Println(resp)
		log.Fatal(err)
		return nil, err
	}
	return resp.BaseUser, nil
}

func FindBaseUserList(author_id []int64) ([]*api.BaseUser, error) {
	rpcReq := &api.DouyinUserFindBaseUserListRequest{
		AuthorId: author_id,
	}
	fmt.Println(rpcReq)
	resp, err := userModelRpcClient.FindBaseUserList(context.Background(), rpcReq)

	if err != nil {
		log.Println(resp)
		log.Fatal(err)
		return nil, err
	}
	return resp.BaseUser, nil
}
