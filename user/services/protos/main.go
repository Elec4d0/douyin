package protos

import (
	"log"
	api "user/services/protos/kitex_gen/api/userservice"
)

func main() {
	svr := api.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
