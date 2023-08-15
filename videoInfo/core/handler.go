package main

import (
	"context"
	api "videoInfo/kitex_gen/api"
)

// VideoInfoProtoBufImpl implements the last service interface defined in the IDL.
type VideoInfoProtoBufImpl struct{}

// GetVideoInfoList implements the VideoInfoProtoBufImpl interface.
func (s *VideoInfoProtoBufImpl) GetVideoInfoList(ctx context.Context, req *api.VideoInfoGetVideoInfoListRequest) (resp *api.VideoInfoGetAuthorVideoInfoListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetVideoInfo implements the VideoInfoProtoBufImpl interface.
func (s *VideoInfoProtoBufImpl) GetVideoInfo(ctx context.Context, req *api.VideoInfoGetVideoInfoRequest) (resp *api.VideoInfoGetVideoInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFeed implements the VideoInfoProtoBufImpl interface.
func (s *VideoInfoProtoBufImpl) GetFeed(ctx context.Context, req *api.VideoInfoGetFeedRequest) (resp *api.VideoInfoGetFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAuthorVideoInfoList implements the VideoInfoProtoBufImpl interface.
func (s *VideoInfoProtoBufImpl) GetAuthorVideoInfoList(ctx context.Context, req *api.VideoInfoGetAuthorVideoInfoListRequest) (resp *api.VideoInfoGetAuthorVideoInfoListResponse, err error) {
	// TODO: Your code here...
	return
}
