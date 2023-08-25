package main

import (
	"comment/comment_deploy/commentsql"
	"comment/sensitiveWord"
	"comment/server/protos"
	api "comment/server/protos/kitex_gen/api/commentserver"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {

	commentsql.MysqlInit()
	sensitiveWord.InitWords()

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:2556")
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewServer(new(protos.CommentServerImpl), server.WithServiceAddr(addr), server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "CommentServer",
		}))
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}

}
