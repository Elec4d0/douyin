package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	protos "videoModel/core"
	api "videoModel/core/kitex_gen/api/videomodelservice"
	model "videoModel/model"
)

func main() {
	//初始化数据库链接
	model.Init()

	//etcd 链接
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(new(protos.VideoModelServiceImpl), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "videoModel"}), server.WithRegistry(r))

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}

}
