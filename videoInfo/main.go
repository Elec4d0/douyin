package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"videoInfo/core"
	api "videoInfo/core/kitex_gen/api/videoinfoservice"
	"videoInfo/rpcApi"
	"videoInfo/tools/redis"
)

func main() {
	//初始化rpcApi链接
	rpcApi.InitRpcClient()

	//初始化redis链接
	redis.InitRedis()

	//etcd 链接
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	//指定IP，对外服务并在ETCD注册
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:13100")
	server := api.NewServer(new(core.VideoInfoServiceImpl), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "videoInfo"}), server.WithRegistry(r))

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}

}
