// Code generated by Kitex v0.6.2. DO NOT EDIT.

package userinfoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	api "userInfo/services/protos/kitex_gen/api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetFullUserInfo(ctx context.Context, Req *api.DouyinUserGetFullUserInfoRequest, callOptions ...callopt.Option) (r *api.DouyinUserGetFullUserInfoResponse, err error)
	GetFullUserInfoList(ctx context.Context, Req *api.DouyinUserGetFullUserInfoListRequest, callOptions ...callopt.Option) (r *api.DouyinUserGetFullUserInfoListResponse, err error)
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
	return &kUserInfoServiceClient{
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

type kUserInfoServiceClient struct {
	*kClient
}

func (p *kUserInfoServiceClient) GetFullUserInfo(ctx context.Context, Req *api.DouyinUserGetFullUserInfoRequest, callOptions ...callopt.Option) (r *api.DouyinUserGetFullUserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFullUserInfo(ctx, Req)
}

func (p *kUserInfoServiceClient) GetFullUserInfoList(ctx context.Context, Req *api.DouyinUserGetFullUserInfoListRequest, callOptions ...callopt.Option) (r *api.DouyinUserGetFullUserInfoListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFullUserInfoList(ctx, Req)
}
