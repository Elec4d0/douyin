package services

import (
	"log"
	"userModel/services/protos"
	api "userModel/services/protos/kitex_gen/api/usermodelservice"
)

func Init() {
	svr := api.NewServer(new(protos.UserModelServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
