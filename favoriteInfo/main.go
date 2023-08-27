package main

import (
	"favoriteInfo/core"
	api "favoriteInfo/core/kitex_gen/api/favoriteinfoservice"
	"favoriteInfo/rpcApi/favoriteModel"
	"favoriteInfo/rpcApi/videoInfo"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	videoInfo.InitVideoInfoRpcClient()
	favoriteModel.InitFavoriteModelRpcClient()

	//etcd 链接
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	//指定IP，对外服务并在ETCD注册
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:14121")
	server := api.NewServer(new(core.FavoriteInfoServiceImpl), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "favoriteInfo"}), server.WithRegistry(r))
	err = server.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
