package main

import (
	"log"
	protos "videoModel/core"
	api "videoModel/core/kitex_gen/api/videomodelprotobuf"
	model "videoModel/model"
)

func main() {
	//初始化数据库链接
	model.Init()
	svr := api.NewServer(new(protos.VideoModelProtoBufImpl))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
