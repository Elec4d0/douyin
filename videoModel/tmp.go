package main

import (
	"time"
	"videoModel/core/kitex_gen/api"
	"videoModel/model"
)

package core

import (
"context"
"time"
api "videoModel/core/kitex_gen/api"
"videoModel/model"
)

// VideoModelProtoBufImpl implements the last service interface defined in the IDL.
type VideoModelProtoBufImpl struct{}

// CreateVideo implements the VideoModelProtoBufImpl interface.
func (s *VideoModelProtoBufImpl) CreateVideo(ctx context.Context, req *api.VideoModelCreateVideoRequest) (resp *api.VideoModelCreateVideoResponse, err error) {
	// TODO: Your code here...

	//写入数据库
	video := &model.Video{
		AuthorID:    req.AuthorId,
		PlayUrl:     req.PlayUrl,
		CoverUrl:    req.CoverUrl,
		Title:       req.Title,
		CreatedTime: time.Now(),
	}
	err = model.CreateVideo(video)

	if err != nil {
		errStr := "数据库创建video对象失败"
		resp = &api.VideoModelCreateVideoResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
		}
		return
	}

	//正常返回
	resp = &api.VideoModelCreateVideoResponse{
		StatusCode: 0,
	}
	return
}

// QueryAuthorWorkCount implements the VideoModelProtoBufImpl interface.
func (s *VideoModelProtoBufImpl) QueryAuthorWorkCount(ctx context.Context, req *api.VideoModelQueryAuthorWorkCountRequest) (resp *api.VideoModelQueryAuthorWorkCountResponse, err error) {
	// TODO: Your code here...
	count, err := model.QueryAuthorWorkCount(req.AuthorId)
	if err != nil {
		errStr := "查询作者作品数失败"
		resp = &api.VideoModelQueryAuthorWorkCountResponse{
			WorkCount:  uint32(count),
			StatusMsg:  &errStr,
			StatusCode: -1,
		}
		return
	}
	resp = &api.VideoModelQueryAuthorWorkCountResponse{
		WorkCount: uint32(count),
	}
	return
}

// QueryAuthorVideoList implements the VideoModelProtoBufImpl interface.
func (s *VideoModelProtoBufImpl) QueryAuthorVideoList(ctx context.Context, req *api.VideoModelQueryAuthorVideoListRequest) (resp *api.VideoModelQueryAuthorVideoListResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryVideoList implements the VideoModelProtoBufImpl interface.
func (s *VideoModelProtoBufImpl) QueryVideoList(ctx context.Context, req *api.VideoModelQueryVideoListRequest) (resp *api.VideoModelQueryVideoListResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryVideo implements the VideoModelProtoBufImpl interface.
func (s *VideoModelProtoBufImpl) QueryVideo(ctx context.Context, req *api.VideoModelQueryVideoRequest) (resp *api.VideoModelQueryVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryVideoFeed implements the VideoModelProtoBufImpl interface.
func (s *VideoModelProtoBufImpl) QueryVideoFeed(ctx context.Context, req *api.VideoModelQueryVideoFeedRequest) (resp *api.VideoModelQueryVideoFeedResponse, err error) {
	// TODO: Your code here...
	return
}
