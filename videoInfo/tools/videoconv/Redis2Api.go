package videoconv

import (
	"videoInfo/core/kitex_gen/api"
	userApi "videoInfo/rpcApi/userInfoAPI/api"
	"videoInfo/tools/redis"
	"videoInfo/tools/userconv"
)

func Redis2Api(redisVideo *redis.Video, isFavorite bool, rpcUser *userApi.FullUser) *api.Video {
	return &api.Video{
		Id:            redisVideo.VideoID,
		Author:        userconv.Rpc2Api(rpcUser),
		PlayUrl:       *redisVideo.PlayUrl,
		CoverUrl:      *redisVideo.CoverUrl,
		Title:         *redisVideo.Title,
		FavoriteCount: redisVideo.FavoriteCount,
		CommentCount:  redisVideo.CommentCount,
		IsFavorite:    isFavorite,
	}
}
