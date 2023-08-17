package core

import (
	"log"
	api "videoInfo/core/kitex_gen/api/videoinfoservice"
)

func main() {
	svr := api.NewServer(new(VideoInfoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
