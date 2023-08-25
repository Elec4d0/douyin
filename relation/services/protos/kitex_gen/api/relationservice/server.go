// Code generated by Kitex v0.7.0. DO NOT EDIT.
package relationservice

import (
	server "github.com/cloudwego/kitex/server"
	api "relation/services/protos/kitex_gen/api"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler api.RelationService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
