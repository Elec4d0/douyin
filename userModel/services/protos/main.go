package protos

import (
	"log"
	api "userModel/services/protos/kitex_gen/api/usermodelservice"
)

func main() {
	svr := api.NewServer(new(UserModelServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
