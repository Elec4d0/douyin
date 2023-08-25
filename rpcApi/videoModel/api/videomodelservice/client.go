// Code generated by Kitex v0.6.2. DO NOT EDIT.

package videomodelservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	api "rpcApi/videoModel/api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateVideo(ctx context.Context, Req *api.VideoModelCreateVideoRequest, callOptions ...callopt.Option) (r *api.VideoModelCreateVideoResponse, err error)
	QueryAuthorWorkCount(ctx context.Context, Req *api.VideoModelQueryAuthorWorkCountRequest, callOptions ...callopt.Option) (r *api.VideoModelQueryAuthorWorkCountResponse, err error)
	QueryAuthorVideoIDList(ctx context.Context, Req *api.VideoModelQueryAuthorVideoIdListRequest, callOptions ...callopt.Option) (r *api.VideoModelQueryAuthorVideoIdListResponse, err error)
	QueryVideoList(ctx context.Context, Req *api.VideoModelQueryVideoListRequest, callOptions ...callopt.Option) (r *api.VideoModelQueryVideoListResponse, err error)
	QueryVideo(ctx context.Context, Req *api.VideoModelQueryVideoRequest, callOptions ...callopt.Option) (r *api.VideoModelQueryVideoResponse, err error)
	QueryVideoFeed(ctx context.Context, Req *api.VideoModelQueryVideoFeedRequest, callOptions ...callopt.Option) (r *api.VideoModelQueryVideoFeedResponse, err error)
	QueryAuthorWorkCountList(ctx context.Context, Req *api.VideoModelQueryAuthorWorkCountListRequest, callOptions ...callopt.Option) (r *api.VideoModelQueryAuthorWorkCountListResponse, err error)
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
	return &kVideoModelServiceClient{
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

type kVideoModelServiceClient struct {
	*kClient
}

func (p *kVideoModelServiceClient) CreateVideo(ctx context.Context, Req *api.VideoModelCreateVideoRequest, callOptions ...callopt.Option) (r *api.VideoModelCreateVideoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateVideo(ctx, Req)
}

func (p *kVideoModelServiceClient) QueryAuthorWorkCount(ctx context.Context, Req *api.VideoModelQueryAuthorWorkCountRequest, callOptions ...callopt.Option) (r *api.VideoModelQueryAuthorWorkCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryAuthorWorkCount(ctx, Req)
}

func (p *kVideoModelServiceClient) QueryAuthorVideoIDList(ctx context.Context, Req *api.VideoModelQueryAuthorVideoIdListRequest, callOptions ...callopt.Option) (r *api.VideoModelQueryAuthorVideoIdListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryAuthorVideoIDList(ctx, Req)
}

func (p *kVideoModelServiceClient) QueryVideoList(ctx context.Context, Req *api.VideoModelQueryVideoListRequest, callOptions ...callopt.Option) (r *api.VideoModelQueryVideoListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryVideoList(ctx, Req)
}

func (p *kVideoModelServiceClient) QueryVideo(ctx context.Context, Req *api.VideoModelQueryVideoRequest, callOptions ...callopt.Option) (r *api.VideoModelQueryVideoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryVideo(ctx, Req)
}

func (p *kVideoModelServiceClient) QueryVideoFeed(ctx context.Context, Req *api.VideoModelQueryVideoFeedRequest, callOptions ...callopt.Option) (r *api.VideoModelQueryVideoFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryVideoFeed(ctx, Req)
}

func (p *kVideoModelServiceClient) QueryAuthorWorkCountList(ctx context.Context, Req *api.VideoModelQueryAuthorWorkCountListRequest, callOptions ...callopt.Option) (r *api.VideoModelQueryAuthorWorkCountListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryAuthorWorkCountList(ctx, Req)
}
