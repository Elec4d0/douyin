package core

import (
	"context"
	"log"
	api "videoInfo/core/kitex_gen/api"
	videoModel "videoInfo/rpcApi/videoModel"
	videoModelApi "videoInfo/rpcApi/videoModel/api"
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
	videoBaseInfoList, _ := videoModel.QueryVideoList(req.VideoIdList)

	//获取用户与视频的喜好关系
	isFavoriteList := getIsFavoriteList(req.UserId, req.VideoIdList)

	//获取视频点赞数
	FavoriteCountList := getFavouriteCountList(req.VideoIdList)

	//获取视频评论数
	CommentCountList := getCommentCountList(req.VideoIdList)

	//非异步项，异步后同步项：获取video作者UserInfo
	var authorIDList []int64
	for _, video := range videoBaseInfoList {
		authorIDList = append(authorIDList, video.AuthorId)
	}
	authorList := getAuthorList(req.UserId, authorIDList)

	resp = &api.VideoInfoGetVideoInfoListResponse{
		VideoList: buildVideoList(videoBaseInfoList, authorList, isFavoriteList, FavoriteCountList, CommentCountList),
	}
	return
}

func buildVideoList(videoBaseInfoList []*videoModelApi.VideoBaseInfo, AuthorList []*api.User, isFavoriteList []bool, favoriteCountList []int64, commentCountList []int64) []*api.Video {
	var videoList []*api.Video

	for i, videoBaseInfo := range videoBaseInfoList {
		var video *api.Video
		//如果基本信息没查到，那么视频也就播放不了，其他信息就没有意义，直接返回nil
		if videoBaseInfo != nil {
			video = &api.Video{
				Id:            videoBaseInfo.VideoId,
				Author:        AuthorList[i],
				PlayUrl:       videoBaseInfo.PlayUrl,
				CoverUrl:      videoBaseInfo.CoverUrl,
				Title:         videoBaseInfo.Title,
				FavoriteCount: favoriteCountList[i],
				CommentCount:  commentCountList[i],
				IsFavorite:    isFavoriteList[i],
			}
		} else {
			video = nil
		}

		videoList = append(videoList, video)
	}
	return videoList
}

func getAuthorList(userID int64, authorIDList []int64) []*api.User {
	return nil
}

func getCommentCountList(videoIDList []int64) []int64 {
	return nil
}

func getFavouriteCountList(videoIDList []int64) []int64 {
	return nil
}

func getIsFavoriteList(userID int64, videoIDList []int64) []bool {
	return nil
}

// GetVideoInfo implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetVideoInfo(ctx context.Context, req *api.VideoInfoGetVideoInfoRequest) (resp *api.VideoInfoGetVideoInfoResponse, err error) {
	// TODO: Your code here...
	resp = &api.VideoInfoGetVideoInfoResponse{
		Video: nil,
	}

	videoBaseInfo, err := videoModel.QueryVideo(req.VideoId)
	if err != nil {
		log.Fatal(err)
		return
	}

	video := &api.Video{
		Id:            req.VideoId,
		Author:        getUserById(videoBaseInfo.AuthorId),
		PlayUrl:       videoBaseInfo.PlayUrl,
		CoverUrl:      videoBaseInfo.CoverUrl,
		Title:         videoBaseInfo.Title,
		FavoriteCount: getFavoriteCount(req.VideoId),
		CommentCount:  getCommentCount(req.VideoId),
		IsFavorite:    getIsFavortite(req.UserId, videoBaseInfo.AuthorId),
	}

	resp = &api.VideoInfoGetVideoInfoResponse{
		Video: video,
	}
	return
}

func getUserById(uid int64) *api.User {
	return nil
}

func getFavoriteCount(videoID int64) int64 {
	return 0
}
func getCommentCount(videoID int64) int64 {
	return 0
}
func getIsFavortite(userID int64, authorID int64) bool {
	return false
}

// GetFeed implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetFeed(ctx context.Context, req *api.VideoInfoGetFeedRequest) (resp *api.VideoInfoGetFeedResponse, err error) {
	// TODO: Your code here...
	//获取视频基本模型
	videoBaseInfoList, _ := videoModel.QueryVideoFeed(req.NextTime)

	var videoIdList []int64
	for _, videoBaseInfo := range videoBaseInfoList {
		videoIdList = append(videoIdList, videoBaseInfo.VideoId)
	}

	//获取用户与视频的喜好关系
	isFavoriteList := getIsFavoriteList(req.UserId, videoIdList)

	//获取视频点赞数
	FavoriteCountList := getFavouriteCountList(videoIdList)

	//获取视频评论数
	CommentCountList := getCommentCountList(videoIdList)

	//非异步项，异步后同步项：获取video作者UserInfo
	var authorIDList []int64
	for _, video := range videoBaseInfoList {
		authorIDList = append(authorIDList, video.AuthorId)
	}
	authorList := getAuthorList(req.UserId, authorIDList)

	resp = &api.VideoInfoGetFeedResponse{
		VideoList: buildVideoList(videoBaseInfoList, authorList, isFavoriteList, FavoriteCountList, CommentCountList),
	}
	return
}

// GetAuthorVideoInfoList implements the VideoInfoServiceImpl interface.
func (s *VideoInfoServiceImpl) GetAuthorVideoInfoList(ctx context.Context, req *api.VideoInfoGetAuthorVideoInfoListRequest) (resp *api.VideoInfoGetAuthorVideoInfoListResponse, err error) {
	// TODO: Your code here...
	//获取视频基本模型
	videoBaseInfoList, _ := videoModel.QueryAuthorVideoList(req.AuthorId)

	var videoIdList []int64
	for _, videoBaseInfo := range videoBaseInfoList {
		videoIdList = append(videoIdList, videoBaseInfo.VideoId)
	}

	//获取用户与视频的喜好关系
	isFavoriteList := getIsFavoriteList(req.UserId, videoIdList)

	//获取视频点赞数
	FavoriteCountList := getFavouriteCountList(videoIdList)

	//获取视频评论数
	CommentCountList := getCommentCountList(videoIdList)

	//非异步项，异步后同步项：获取video作者UserInfo
	var authorIDList []int64
	for _, video := range videoBaseInfoList {
		authorIDList = append(authorIDList, video.AuthorId)
	}
	authorList := getAuthorList(req.UserId, authorIDList)

	resp = &api.VideoInfoGetAuthorVideoInfoListResponse{
		VideoList: buildVideoList(videoBaseInfoList, authorList, isFavoriteList, FavoriteCountList, CommentCountList),
	}
	return
}
