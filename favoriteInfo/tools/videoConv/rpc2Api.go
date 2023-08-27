package videoConv

import (
	"favoriteInfo/core/kitex_gen/api"
	videoInfo "favoriteInfo/rpcApi/videoInfo/api"
	"favoriteInfo/tools/userConv"
)

func Rpc2Api(video *videoInfo.Video) *api.FavoriteVideo {
	return &api.FavoriteVideo{
		Id:            video.Id,
		Author:        userConv.Rpc2Api(video.Author),
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite,
		Title:         video.Title,
	}
}

func BatchRpc2Api(rpcVideoList []*videoInfo.Video) []*api.FavoriteVideo {
	var apiVideoList = make([]*api.FavoriteVideo, len(rpcVideoList))

	for i, rpcVideo := range rpcVideoList {
		apiVideoList[i] = Rpc2Api(rpcVideo)
	}
	return apiVideoList
}
