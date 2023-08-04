package main

import (
	"user/model"
	"user/services"
)

func main() {
	// init database
	model.Init()
	// init rpc services
	services.Init()
}
