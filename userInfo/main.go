package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"userInfo/services/protos"
	api "userInfo/services/protos/kitex_gen/api/userinfoservice"
	userModelServices "userInfo/userModelAPI"
	"userInfo/videoModel"
)

func main() {
	//初始化rpcApi链接
	userModelServices.InitUserModelRpcClient()
	videoModel.InitVideoModelRpcClient()

	//etcd 链接
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	//指定IP，对外服务并在ETCD注册
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:15100")
	server := api.NewServer(new(protos.UserInfoServiceImpl), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "userInfoService"}), server.WithRegistry(r))

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
