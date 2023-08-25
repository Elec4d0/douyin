package core

import (
	"context"
	api "videoInfo/core/kitex_gen/api"
	"videoInfo/rpcApi"
	"videoInfo/tools/query"
	"videoInfo/tools/redis"
	"videoInfo/tools/userconv"
)

// VideoInfoServiceImpl implements the last service interface defined in the IDL.
type VideoInfoServiceImpl struct{}

// GetVideoInfoList implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetVideoInfoList(ctx context.Context, req *api.VideoInfoGetVideoInfoListRequest) (resp *api.VideoInfoGetVideoInfoListResponse, err error) {
	// TODO: Your code here...
	//检查每一个视频的缓存情况
	isVideoListCache, cacheType := redis.CheckAllVideoCache(req.VideoIdList)

	var queryOK bool
	var videoList []*api.Video

	if cacheType == 1 {
		//所有视频均在缓存，仅查redis
		videoList, queryOK = query.RedisQueryVideoList(req.UserId, req.VideoIdList)
	} else if cacheType == 0 {
		//部分视频缓存，混合查询
		videoList, queryOK = query.MixQueryVideoList(req.UserId, req.VideoIdList, isVideoListCache)
	}

	if queryOK {
		resp = &api.VideoInfoGetVideoInfoListResponse{
			VideoList: videoList,
		}
		return
	}

	//所有视频均未缓存, cacheType = -1
	//上面的两个if查询错误均回落于此
	videoList = query.RpcQueryVideoList(req.UserId, req.VideoIdList)
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
	query.CacheVideo(video)

	resp = &api.VideoInfoGetVideoInfoResponse{Video: video}
	return
}

// GetFeed implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetFeed(ctx context.Context, req *api.VideoInfoGetFeedRequest) (resp *api.VideoInfoGetFeedResponse, err error) {
	// TODO: Your code here...
	if redis.CheckFeedExists() {
		//获取Feed流ID列表
		videoIDList := redis.QueryFeedIDList()
		//检查每一个视频的缓存情况
		isVideoListCache, cacheType := redis.CheckAllVideoCache(videoIDList)

		//执行缓存查询
		var queryOK bool
		var videoList []*api.Video
		var str string
		if cacheType == 1 {
			//所有视频均在缓存，仅查redis
			videoList, queryOK = query.RedisQueryVideoList(req.UserId, videoIDList)
			str = "本次查询缓存全命中"
		} else if cacheType == 0 {
			//部分视频缓存，混合查询
			videoList, queryOK = query.MixQueryVideoList(req.UserId, videoIDList, isVideoListCache)
			str = "本次查询缓存半命中"
		}

		if queryOK {
			resp = &api.VideoInfoGetFeedResponse{
				StatusMsg: &str,
				VideoList: videoList,
			}
			return
		}

		str = "本次查询缓存未命中, 但已缓存Feed列表"
		videoList = query.RpcQueryVideoList(req.UserId, videoIDList)
		resp = &api.VideoInfoGetFeedResponse{
			StatusMsg: &str,
			VideoList: videoList,
		}
		return
	}

	//获取视频基本模型
	videoBaseInfoList := rpcApi.QueryVideoFeed(req.NextTime)
	var videoIdList []int64
	for _, videoBaseInfo := range videoBaseInfoList {
		videoIdList = append(videoIdList, videoBaseInfo.VideoId)
	}

	//获取用户与视频的喜好关系
	isFavoriteList := rpcApi.GetIsFavoriteList(req.UserId, videoIdList)
	//获取视频点赞数
	FavoriteCountList := rpcApi.GetFavouriteCountList(videoIdList)
	//获取视频评论数
	CommentCountList := rpcApi.GetCommentCountList(videoIdList)

	//非异步项，异步后同步项：获取video作者UserInfo
	var authorIDList []int64
	for _, video := range videoBaseInfoList {
		authorIDList = append(authorIDList, video.AuthorId)
	}
	rpcAuthorList := rpcApi.GetAuthorList(req.UserId, authorIDList)
	apiAuthorList := userconv.BatchRpc2Api(rpcAuthorList)
	videoList := rpcApi.BuildVideoList(videoBaseInfoList, apiAuthorList, isFavoriteList, FavoriteCountList, CommentCountList)

	//这里做defer实现异步
	query.CacheVideoList(videoList)
	redis.CacheFeed(videoIdList)
	str := "本次查询缓存未命中"
	resp = &api.VideoInfoGetFeedResponse{
		VideoList: videoList,
		StatusMsg: &str,
	}
	return
}

// GetAuthorVideoInfoList implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetAuthorVideoInfoList(ctx context.Context, req *api.VideoInfoGetAuthorVideoInfoListRequest) (resp *api.VideoInfoGetAuthorVideoInfoListResponse, err error) {
	// TODO: Your code here...
	//redis 缓存过用户作品列表的videoID
	if redis.CheckAuthorVideoIDListExists(req.AuthorId) {
		//从redis中拿出缓存的用户作品列表
		videoIDList := redis.QueryAuthorVideoIDList(req.AuthorId)
		//检查每一个视频的缓存情况
		isVideoListCache, cacheType := redis.CheckAllVideoCache(videoIDList)

		var queryOK bool
		var videoList []*api.Video

		if cacheType == 1 {
			//所有视频均在缓存，仅查redis
			videoList, queryOK = query.RedisQueryVideoList(req.UserId, videoIDList)
		} else if cacheType == 0 {
			//部分视频缓存，混合查询
			videoList, queryOK = query.MixQueryVideoList(req.UserId, videoIDList, isVideoListCache)
		}

		if queryOK {
			resp = &api.VideoInfoGetAuthorVideoInfoListResponse{
				VideoList: videoList,
			}
			return
		}

		//所有视频均未缓存, cacheType = -1
		//上面的两个if查询错误均回落于此
		videoList = query.RpcQueryVideoList(req.UserId, videoIDList)
		resp = &api.VideoInfoGetAuthorVideoInfoListResponse{
			VideoList: videoList,
		}
		return
	}

	//作者视频列表未缓存

	//获取视频基本模型
	videoBaseInfoList := rpcApi.QueryAuthorVideoList(req.AuthorId)
	var videoIdList []int64
	for _, videoBaseInfo := range videoBaseInfoList {
		videoIdList = append(videoIdList, videoBaseInfo.VideoId)
	}

	//获取用户与视频的喜好关系
	isFavoriteList := rpcApi.GetIsFavoriteList(req.UserId, videoIdList)
	//获取视频点赞数
	FavoriteCountList := rpcApi.GetFavouriteCountList(videoIdList)
	//获取视频评论数
	CommentCountList := rpcApi.GetCommentCountList(videoIdList)

	//非异步项，异步后同步项：获取video作者UserInfo
	var authorIDList []int64
	for _, video := range videoBaseInfoList {
		authorIDList = append(authorIDList, video.AuthorId)
	}
	rpcAuthorList := rpcApi.GetAuthorList(req.UserId, authorIDList)
	apiAuthorList := userconv.BatchRpc2Api(rpcAuthorList)
	videoList := rpcApi.BuildVideoList(videoBaseInfoList, apiAuthorList, isFavoriteList, FavoriteCountList, CommentCountList)

	query.CacheVideoList(videoList)
	redis.CacheAuthorVideoIDList(req.AuthorId, videoIdList)
	resp = &api.VideoInfoGetAuthorVideoInfoListResponse{
		VideoList: videoList,
	}
	return
}
