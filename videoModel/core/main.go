package core

import (
	"log"
	api "videoModel/core/kitex_gen/api/videomodelprotobuf"
)

func main() {
	svr := api.NewServer(new(VideoModelProtoBufImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
