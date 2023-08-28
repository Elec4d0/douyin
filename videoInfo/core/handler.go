package core

import (
	"context"
	api "videoInfo/core/kitex_gen/api"
	"videoInfo/tools/query"
	"videoInfo/tools/redis"
	"videoInfo/tools/videoconv"
)

// VideoInfoServiceImpl implements the last service interface defined in the IDL.
type VideoInfoServiceImpl struct{}

// GetVideoInfoList implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetVideoInfoList(ctx context.Context, req *api.VideoInfoGetVideoInfoListRequest) (resp *api.VideoInfoGetVideoInfoListResponse, err error) {
	// TODO: Your code here...
	//检查每一个视频的缓存情况
	videoList := query.GetVideoInfoList(req.UserId, req.VideoIdList)
	resp = &api.VideoInfoGetVideoInfoListResponse{
		VideoList: videoList,
	}
	return
}

// GetVideoInfo implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetVideoInfo(ctx context.Context, req *api.VideoInfoGetVideoInfoRequest) (resp *api.VideoInfoGetVideoInfoResponse, err error) {
	// TODO: Your code here...

	var video *api.Video
	var cacheQueryOK bool
	//缓存是否命中
	if redis.CheckVideoExists(req.VideoId) {
		video, cacheQueryOK = query.RedisQueryVideo(req.UserId, req.VideoId)
	}
	//缓存查询成功则返回
	if cacheQueryOK {
		resp = &api.VideoInfoGetVideoInfoResponse{Video: video}
		return
	}

	//缓存失效或缓存查询失败，回落到持久化层进行查询
	video = query.RpcQueryVideo(req.UserId, req.VideoId)

	//缓存该video，这里做defer
	defer redis.CacheVideo(videoconv.Api2Redis(video))

	resp = &api.VideoInfoGetVideoInfoResponse{Video: video}
	return
}

// GetFeed implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetFeed(ctx context.Context, req *api.VideoInfoGetFeedRequest) (resp *api.VideoInfoGetFeedResponse, err error) {
	// TODO: Your code here...
	videoList, nextTime := query.GetFeed(req.UserId, req.NextTime)
	resp = &api.VideoInfoGetFeedResponse{
		VideoList: videoList,
		NextTime:  nextTime,
	}
	return
}

// GetAuthorVideoInfoList implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetAuthorVideoInfoList(ctx context.Context, req *api.VideoInfoGetAuthorVideoInfoListRequest) (resp *api.VideoInfoGetAuthorVideoInfoListResponse, err error) {
	// TODO: Your code here...
	videoList := query.GetAuthorVideoInfoList(req.UserId, req.AuthorId)
	resp = &api.VideoInfoGetAuthorVideoInfoListResponse{
		VideoList: videoList,
	}
	return
}
