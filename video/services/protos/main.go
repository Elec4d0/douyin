package protos

import (
	"log"
	api "video/services/protos/kitex_gen/api/feedprotobuf"
)

func main() {
	svr := api.NewServer(new(FeedProtoBufImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
