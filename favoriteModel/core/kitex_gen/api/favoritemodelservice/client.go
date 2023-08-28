// Code generated by Kitex v0.6.2. DO NOT EDIT.

package favoritemodelservice

import (
	"context"
	api "favoriteModel/core/kitex_gen/api"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	QueryFavoriteList(ctx context.Context, Req *api.FavoriteModelQueryFavoriteListRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryFavoriteListResponse, err error)
	FavoriteAction(ctx context.Context, Req *api.FavoriteModelFavoriteActionRequest, callOptions ...callopt.Option) (r *api.FavoriteModelFavoriteActionResponse, err error)
	QueryUserFavoriteCount(ctx context.Context, Req *api.FavoriteModelQueryUserFavoriteCountRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryUserFavoriteCountResponse, err error)
	QueryUserFavoriteCountList(ctx context.Context, Req *api.FavoriteModelQueryUserFavoriteCountListRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryUserFavoriteCountListResponse, err error)
	QueryVideoFavoriteCountList(ctx context.Context, Req *api.FavoriteModelQueryVideoFavoriteCountListRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryVideoFavoriteCountListResponse, err error)
	QueryVideoFavoriteCount(ctx context.Context, Req *api.FavoriteModelQueryVideoFavoriteCountRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryVideoFavoriteCountResponse, err error)
	QueryIsUserFavoriteList(ctx context.Context, Req *api.FavoriteModelQueryIsUserFavoriteListRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryIsUserFavoriteListResponse, err error)
	QueryIsUserFavorite(ctx context.Context, Req *api.FavoriteModelQueryIsUserFavoriteRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryIsUserFavoriteResponse, err error)
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
	return &kFavoriteModelServiceClient{
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

type kFavoriteModelServiceClient struct {
	*kClient
}

func (p *kFavoriteModelServiceClient) QueryFavoriteList(ctx context.Context, Req *api.FavoriteModelQueryFavoriteListRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryFavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryFavoriteList(ctx, Req)
}

func (p *kFavoriteModelServiceClient) FavoriteAction(ctx context.Context, Req *api.FavoriteModelFavoriteActionRequest, callOptions ...callopt.Option) (r *api.FavoriteModelFavoriteActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteAction(ctx, Req)
}

func (p *kFavoriteModelServiceClient) QueryUserFavoriteCount(ctx context.Context, Req *api.FavoriteModelQueryUserFavoriteCountRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryUserFavoriteCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryUserFavoriteCount(ctx, Req)
}

func (p *kFavoriteModelServiceClient) QueryUserFavoriteCountList(ctx context.Context, Req *api.FavoriteModelQueryUserFavoriteCountListRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryUserFavoriteCountListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryUserFavoriteCountList(ctx, Req)
}

func (p *kFavoriteModelServiceClient) QueryVideoFavoriteCountList(ctx context.Context, Req *api.FavoriteModelQueryVideoFavoriteCountListRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryVideoFavoriteCountListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryVideoFavoriteCountList(ctx, Req)
}

func (p *kFavoriteModelServiceClient) QueryVideoFavoriteCount(ctx context.Context, Req *api.FavoriteModelQueryVideoFavoriteCountRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryVideoFavoriteCountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryVideoFavoriteCount(ctx, Req)
}

func (p *kFavoriteModelServiceClient) QueryIsUserFavoriteList(ctx context.Context, Req *api.FavoriteModelQueryIsUserFavoriteListRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryIsUserFavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryIsUserFavoriteList(ctx, Req)
}

func (p *kFavoriteModelServiceClient) QueryIsUserFavorite(ctx context.Context, Req *api.FavoriteModelQueryIsUserFavoriteRequest, callOptions ...callopt.Option) (r *api.FavoriteModelQueryIsUserFavoriteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryIsUserFavorite(ctx, Req)
}
