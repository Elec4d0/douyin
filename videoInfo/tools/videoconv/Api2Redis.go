package videoconv

import (
	"videoInfo/core/kitex_gen/api"
	"videoInfo/tools/redis"
)

func Api2Redis(apiVideo *api.Video) *redis.Video {
	if apiVideo == nil {
		return nil
	}
	return &redis.Video{
		VideoID:       apiVideo.Id,
		AuthorID:      apiVideo.Author.Id,
		PlayUrl:       &apiVideo.PlayUrl,
		CoverUrl:      &apiVideo.CoverUrl,
		FavoriteCount: apiVideo.FavoriteCount,
		CommentCount:  apiVideo.CommentCount,
		Title:         &apiVideo.Title,
	}
}

func BatchApi2Redis(apiVideoList []*api.Video) []*redis.Video {
	var redisVideoList = make([]*redis.Video, len(apiVideoList))
	for i, apiVideo := range apiVideoList {
		redisVideoList[i] = Api2Redis(apiVideo)
	}
	return redisVideoList
}
