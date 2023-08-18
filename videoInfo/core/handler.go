package core

import (
	"context"
	"log"
	api "videoInfo/core/kitex_gen/api"
	"videoInfo/rpcApi"
)

// VideoInfoServiceImpl implements the last service interface defined in the IDL.
type VideoInfoServiceImpl struct{}

// GetVideoInfoList implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetVideoInfoList(ctx context.Context, req *api.VideoInfoGetVideoInfoListRequest) (resp *api.VideoInfoGetVideoInfoListResponse, err error) {
	// TODO: Your code here...
	resp = &api.VideoInfoGetVideoInfoListResponse{
		VideoList: nil,
	}

	//获取Video基本模型
	videoBaseInfoList := rpcApi.QueryVideoList(req.VideoIdList)

	//获取用户与视频的喜好关系
	isFavoriteList := rpcApi.GetIsFavoriteList(req.UserId, req.VideoIdList)

	//获取视频点赞数
	FavoriteCountList := rpcApi.GetFavouriteCountList(req.VideoIdList)

	//获取视频评论数
	CommentCountList := rpcApi.GetCommentCountList(req.VideoIdList)

	//非异步项，异步后同步项：获取video作者UserInfo
	var authorIDList []int64
	for _, video := range videoBaseInfoList {
		authorIDList = append(authorIDList, video.AuthorId)
	}
	authorList := rpcApi.GetAuthorList(req.UserId, authorIDList)

	resp = &api.VideoInfoGetVideoInfoListResponse{
		VideoList: rpcApi.BuildVideoList(videoBaseInfoList, authorList, isFavoriteList, FavoriteCountList, CommentCountList),
	}
	return
}

// GetVideoInfo implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetVideoInfo(ctx context.Context, req *api.VideoInfoGetVideoInfoRequest) (resp *api.VideoInfoGetVideoInfoResponse, err error) {
	// TODO: Your code here...
	resp = &api.VideoInfoGetVideoInfoResponse{
		Video: nil,
	}

	videoBaseInfo := rpcApi.QueryVideo(req.VideoId)
	if err != nil {
		log.Fatal(err)
		return
	}

	video := &api.Video{
		Id:            req.VideoId,
		Author:        rpcApi.GetUserById(req.UserId, videoBaseInfo.AuthorId),
		PlayUrl:       videoBaseInfo.PlayUrl,
		CoverUrl:      videoBaseInfo.CoverUrl,
		Title:         videoBaseInfo.Title,
		FavoriteCount: rpcApi.GetFavoriteCount(req.VideoId),
		CommentCount:  rpcApi.GetCommentCount(req.VideoId),
		IsFavorite:    rpcApi.GetIsFavortite(req.UserId, videoBaseInfo.AuthorId),
	}

	resp = &api.VideoInfoGetVideoInfoResponse{
		Video: video,
	}
	return
}

// GetFeed implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetFeed(ctx context.Context, req *api.VideoInfoGetFeedRequest) (resp *api.VideoInfoGetFeedResponse, err error) {
	// TODO: Your code here...
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
	authorList := rpcApi.GetAuthorList(req.UserId, authorIDList)

	resp = &api.VideoInfoGetFeedResponse{
		VideoList: rpcApi.BuildVideoList(videoBaseInfoList, authorList, isFavoriteList, FavoriteCountList, CommentCountList),
	}
	return
}

// GetAuthorVideoInfoList implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetAuthorVideoInfoList(ctx context.Context, req *api.VideoInfoGetAuthorVideoInfoListRequest) (resp *api.VideoInfoGetAuthorVideoInfoListResponse, err error) {
	// TODO: Your code here...
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
	authorList := rpcApi.GetAuthorList(req.UserId, authorIDList)

	resp = &api.VideoInfoGetAuthorVideoInfoListResponse{
		VideoList: rpcApi.BuildVideoList(videoBaseInfoList, authorList, isFavoriteList, FavoriteCountList, CommentCountList),
	}
	return
}
