package videoconv

import (
	"videoInfo/core/kitex_gen/api"
	userApi "videoInfo/rpcApi/userInfoAPI/api"
	videoModelApi "videoInfo/rpcApi/videoModel/api"
	"videoInfo/tools/userconv"
)

func Rpc2Api(rpcVideoModel *videoModelApi.VideoBaseInfo, rpcUser *userApi.FullUser, favoriteCount int64, commentCount int64, isFavorite bool) *api.Video {
	if rpcVideoModel == nil {
		return nil
	}
	return &api.Video{
		Id:            rpcVideoModel.VideoId,
		Author:        userconv.Rpc2Api(rpcUser),
		PlayUrl:       rpcVideoModel.PlayUrl,
		CoverUrl:      rpcVideoModel.CoverUrl,
		Title:         rpcVideoModel.Title,
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		IsFavorite:    isFavorite,
	}
}

func BatchRpc2Api(rpcVideoModelList []*videoModelApi.VideoBaseInfo, rpcAuthorList []*userApi.FullUser, favoriteCountList []int64, commentCountList []int64, isFavoriteList []bool) []*api.Video {
	var videoList = make([]*api.Video, len(rpcVideoModelList))
	apiAuthorList := userconv.BatchRpc2Api(rpcAuthorList)

	for i, _ := range rpcVideoModelList {
		videoList[i] = buildApiVideo(rpcVideoModelList[i], apiAuthorList[i], favoriteCountList[i], commentCountList[i], isFavoriteList[i])
	}
	return videoList
}

func buildApiVideo(rpcVideoModel *videoModelApi.VideoBaseInfo, apiUser *api.User, favoriteCount int64, commentCount int64, isFavorite bool) *api.Video {
	//如果基本信息没查到，那么视频也就播放不了，其他信息就没有意义，直接赋予nil
	if rpcVideoModel == nil {
		return nil
	}
	return &api.Video{
		Id:            rpcVideoModel.VideoId,
		Author:        apiUser,
		PlayUrl:       rpcVideoModel.PlayUrl,
		CoverUrl:      rpcVideoModel.CoverUrl,
		Title:         rpcVideoModel.Title,
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		IsFavorite:    isFavorite,
	}
}
