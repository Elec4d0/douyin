// Code generated by Kitex v0.6.2. DO NOT EDIT.

package commentserver

import (
	api "comment/server/protos/kitex_gen/api"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CommentAction(ctx context.Context, Req *api.DouyinCommentActionRequest, callOptions ...callopt.Option) (r *api.DouyinCommentActionResponse, err error)
	CommentList(ctx context.Context, Req *api.DouyinCommentListRequest, callOptions ...callopt.Option) (r *api.DouyinCommentListResponse, err error)
	CommentCount(ctx context.Context, Req *api.DouyinCommentserverCommentcountRequest, callOptions ...callopt.Option) (r *api.DouyinCommentserverCommentcountResponse, err error)
	CommentAllCount(ctx context.Context, Req *api.DouyinCommentserverCommentallcountRequest, callOptions ...callopt.Option) (r *api.DouyinCommentserverCommentallcountResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCommentServerClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCommentServerClient struct {
	*kClient
}

func (p *kCommentServerClient) CommentAction(ctx context.Context, Req *api.DouyinCommentActionRequest, callOptions ...callopt.Option) (r *api.DouyinCommentActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentAction(ctx, Req)
}

func (p *kCommentServerClient) CommentList(ctx context.Context, Req *api.DouyinCommentListRequest, callOptions ...callopt.Option) (r *api.DouyinCommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, Req)
}

func (p *kCommentServerClient) CommentCount(ctx context.Context, Req *api.DouyinCommentserverCommentcountRequest, callOptions ...callopt.Option) (r *api.DouyinCommentserverCommentcountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentCount(ctx, Req)
}

func (p *kCommentServerClient) CommentAllCount(ctx context.Context, Req *api.DouyinCommentserverCommentallcountRequest, callOptions ...callopt.Option) (r *api.DouyinCommentserverCommentallcountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentAllCount(ctx, Req)
}
