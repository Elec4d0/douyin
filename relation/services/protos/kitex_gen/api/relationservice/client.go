// Code generated by Kitex v0.7.0. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	api "relation/services/protos/kitex_gen/api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RelationAction(ctx context.Context, Req *api.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *api.DouyinRelationActionResponse, err error)
	RelationFollowList(ctx context.Context, Req *api.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *api.DouyinRelationFollowListResponse, err error)
	RelationFollowerList(ctx context.Context, Req *api.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *api.DouyinRelationFollowerListResponse, err error)
	RelationFriendList(ctx context.Context, Req *api.DouyinRelationFriendListRequest, callOptions ...callopt.Option) (r *api.DouyinRelationFriendListResponse, err error)
	GetOneRelation(ctx context.Context, Req *api.DouyinRelationSearchRequest, callOptions ...callopt.Option) (r *api.DouyinRelationSearchResponse, err error)
	GetListRelation(ctx context.Context, Req *api.DouyinRelationSearchListRequest, callOptions ...callopt.Option) (r *api.DouyinRelationSearchListResponse, err error)
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
	return &kRelationServiceClient{
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

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) RelationAction(ctx context.Context, Req *api.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *api.DouyinRelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, Req)
}

func (p *kRelationServiceClient) RelationFollowList(ctx context.Context, Req *api.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *api.DouyinRelationFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationFollowList(ctx, Req)
}

func (p *kRelationServiceClient) RelationFollowerList(ctx context.Context, Req *api.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *api.DouyinRelationFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationFollowerList(ctx, Req)
}

func (p *kRelationServiceClient) RelationFriendList(ctx context.Context, Req *api.DouyinRelationFriendListRequest, callOptions ...callopt.Option) (r *api.DouyinRelationFriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationFriendList(ctx, Req)
}

func (p *kRelationServiceClient) GetOneRelation(ctx context.Context, Req *api.DouyinRelationSearchRequest, callOptions ...callopt.Option) (r *api.DouyinRelationSearchResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetOneRelation(ctx, Req)
}

func (p *kRelationServiceClient) GetListRelation(ctx context.Context, Req *api.DouyinRelationSearchListRequest, callOptions ...callopt.Option) (r *api.DouyinRelationSearchListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetListRelation(ctx, Req)
}
