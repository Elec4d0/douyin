// Code generated by Kitex v0.6.2. DO NOT EDIT.
package feedprotobuf

import (
	server "github.com/cloudwego/kitex/server"
	api "video/services/protos/kitex_gen/api"
)

// NewServer creates a services.Server with the given handler and options.
func NewServer(handler api.FeedProtoBuf, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
