package main

import (
	"log"
	api "relation/services/protos/kitex_gen/api/relationservice"
)

func main() {
	svr := api.NewServer(new(RelationServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
