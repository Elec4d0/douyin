// Code generated by Kitex v0.6.2. DO NOT EDIT.

package feedprotobuf

import (
	api "gateway/microService/feed/api"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a services.Invoker with the given handler and options.
func NewInvoker(handler api.FeedProtoBuf, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
