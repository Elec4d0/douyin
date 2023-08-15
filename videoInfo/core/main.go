package main

import (
	"log"
	api "videoInfo/kitex_gen/api/videoinfoprotobuf"
)

func main() {
	svr := api.NewServer(new(VideoInfoProtoBufImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
