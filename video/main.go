package main

import (
	"log"
	model "video/model"
	"video/rpcClient"
	protos "video/services/protos"
	api "video/services/protos/kitex_gen/api/feedprotobuf"
	oss "video/tools/oss"
)

func main() {
	//初始化数据库链接
	model.Init()
	oss.InitOss()
	rpcClient.InitUserRpcClient()

	svr := api.NewServer(new(protos.FeedProtoBufImpl))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
