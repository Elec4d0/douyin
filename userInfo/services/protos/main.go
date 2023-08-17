package protos

import (
	"log"
	api "userInfo/services/protos/kitex_gen/api/userinfoservice"
)

func main() {
	svr := api.NewServer(new(UserInfoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
