package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	"videoPublish/core"
	api "videoPublish/core/kitex_gen/api/videopublishservice"
	"videoPublish/rpcApi/videoModel"
)

func main() {
	//初始化rpcApi链接
	videoModel.InitVideoModelRpcClient()

	//etcd链接
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	//指定IP，对外服务并在ETCD注册
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:13101")
	svr := api.NewServer(new(core.VideoPublishServiceImpl), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "videoPublish"}), server.WithRegistry(r))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
