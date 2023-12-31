package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"user/services/protos"
	api "user/services/protos/kitex_gen/api/userservice"
	userModelServices "user/userModelAPI"
)

func main() {
	//初始化rpcApi链接
	userModelServices.InitUserModelRpcClient()

	//etcd 链接
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}
	//addr
	//指定IP，对外服务并在ETCD注册
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:15101")
	server := api.NewServer(new(protos.UserServiceImpl), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "userService"}), server.WithRegistry(r))

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
