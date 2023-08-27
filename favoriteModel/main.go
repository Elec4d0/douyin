package main

import (
	"favoriteModel/core"
	api "favoriteModel/core/kitex_gen/api/favoritemodelservice"
	"favoriteModel/model"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	//数据库初始化链接
	model.Init()

	//etcd 链接
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	//指定IP，对外服务并在ETCD注册
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:14120")
	server := api.NewServer(new(core.FavoriteModelServiceImpl), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "favoriteModel"}), server.WithRegistry(r))
	err = server.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
