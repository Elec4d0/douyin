package services

import (
	"fmt"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"user/services/protos"
	api "user/services/protos/kitex_gen/api/userservice"
)

func Init() {
	//addr, _ := net.ResolveTCPAddr("tcp", ":8801")
	//var opts []server.Option
	//opts = append(opts, server.WithServiceAddr(addr))
	addr, err := net.ResolveTCPAddr("tcp", ":8887")
	fmt.Println(addr)
	fmt.Println(err)
	svr := api.NewServer(new(protos.UserServiceImpl), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
