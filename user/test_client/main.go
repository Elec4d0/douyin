package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"log"
	"user/services/protos/kitex_gen/api"
	"user/services/protos/kitex_gen/api/userservice"
)

func main() {
	clientWe, err := userservice.NewClient("UserService", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.DouyinUserRegisterRequest{
		Username: "在山的那边",
		Password: "123456",
	}
	resp, err := clientWe.UserRegister(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

}
