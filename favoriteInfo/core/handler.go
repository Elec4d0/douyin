package core

import (
	"context"
	api "favoriteInfo/core/kitex_gen/api"
	"favoriteInfo/rpcApi/favoriteModel"
	"favoriteInfo/rpcApi/videoInfo"
	"favoriteInfo/tools/videoConv"
	"log"
)

// FavoriteInfoServiceImpl implements the last service interface defined in the IDL.
type FavoriteInfoServiceImpl struct{}

// QueryFavoriteList implements the FavoriteInfoServiceImpl interface.
func (s *FavoriteInfoServiceImpl) QueryFavoriteList(ctx context.Context, req *api.FavoriteInfoQueryFavoriteListRequest) (resp *api.FavoriteInfoQueryFavoriteListResponse, err error) {
	// TODO: Your code here...
	log.Println("查找的用户ID：", req.SearchId)
	videoIDList, err := favoriteModel.QueryFavoriteList(req.SearchId)
	log.Println("从model层拿到的喜欢列表ID", videoIDList)
	if err != nil {
		errStr := "未查询到该用户的喜欢列表"
		resp = &api.FavoriteInfoQueryFavoriteListResponse{
			StatusCode: 0,
			StatusMsg:  &errStr,
			VideoList:  nil,
		}
	}

	videoList, _ := videoInfo.GetVideoInfoList(req.UserId, videoIDList)
	log.Println("videoInfo 查询到的videoList", videoList)
	resp = &api.FavoriteInfoQueryFavoriteListResponse{
		StatusCode: 0,
		VideoList:  videoConv.BatchRpc2Api(videoList),
	}
	return
}

// FavoriteAction implements the FavoriteInfoServiceImpl interface.
func (s *FavoriteInfoServiceImpl) FavoriteAction(ctx context.Context, req *api.FavoriteInfoFavoriteActionRequest) (resp *api.FavoriteInfoFavoriteActionResponse, err error) {
	// TODO: Your code here...
	video, _ := videoInfo.GetVideoInfo(-1, req.VideoId)
	err = favoriteModel.FavoriteAction(req.UserId, req.VideoId, video.Author.Id, req.ActionType)
	if err != nil {
		errStr := "执行点赞失败"
		resp = &api.FavoriteInfoFavoriteActionResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
		}
		return
	}
	resp = &api.FavoriteInfoFavoriteActionResponse{
		StatusCode: 0,
	}
	return
}
