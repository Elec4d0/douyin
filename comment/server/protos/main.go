package main

import (
	"comment/comment_deploy/comment_mysql"
	api "comment/server/protos/kitex_gen/api/commentserver"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	comment_mysql.Init()

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:2556")
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewServer(new(CommentServerImpl), server.WithServiceAddr(addr), server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "CommentServer",
		}))
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}

}
