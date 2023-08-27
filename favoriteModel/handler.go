package core

import (
	"context"
	"errors"
	api "favoriteModel/core/kitex_gen/api"
	"favoriteModel/model"
)

// FavoriteModelServiceImpl implements the last service interface defined in the IDL.
type FavoriteModelServiceImpl struct{}

// QueryFavoriteList implements the FavoriteModelServiceImpl interface.
func (s *FavoriteModelServiceImpl) QueryFavoriteList(ctx context.Context, req *api.FavoriteModelQueryFavoriteListRequest) (resp *api.FavoriteModelQueryFavoriteListResponse, err error) {
	// TODO: Your code here...
	videoIDList, err := model.QueryUserFavoriteList(req.UserId)
	if err != nil {
		errStr := "FavoriteList查询失败"
		resp = &api.FavoriteModelQueryFavoriteListResponse{
			StatusCode:  -1,
			StatusMsg:   &errStr,
			VideoIdList: nil,
		}
	}

	resp = &api.FavoriteModelQueryFavoriteListResponse{
		VideoIdList: videoIDList,
	}

	return
}

// FavoriteAction implements the FavoriteModelServiceImpl interface.
func (s *FavoriteModelServiceImpl) FavoriteAction(ctx context.Context, req *api.FavoriteModelFavoriteActionRequest) (resp *api.FavoriteModelFavoriteActionResponse, err error) {
	// TODO: Your code here...
	resp = new(api.FavoriteModelFavoriteActionResponse)
	if req.ActionType == 1 {
		err = model.CreateLikeVideo(req.UserId, req.VideoId, req.AuthorId)
	} else if req.ActionType == 2 {
		err = model.RemoveLikeVideo(req.UserId, req.VideoId, req.AuthorId)
	} else {
		err = errors.New("无效请求类型")
	}

	if err != nil {
		errStr := "操作失败"
		resp.StatusMsg = &errStr
		resp.StatusCode = 1
		return
	}

	return
}

// QueryUserFavoriteCount implements the FavoriteModelServiceImpl interface.
func (s *FavoriteModelServiceImpl) QueryUserFavoriteCount(ctx context.Context, req *api.FavoriteModelQueryUserFavoriteCountRequest) (resp *api.FavoriteModelQueryUserFavoriteCountResponse, err error) {
	// TODO: Your code here...
	likeCount, favoritedCount, _ := model.QueryUserFavoriteCount(req.UserId)
	resp = &api.FavoriteModelQueryUserFavoriteCountResponse{}
	return
}

// QueryUserFavoriteCountList implements the FavoriteModelServiceImpl interface.
func (s *FavoriteModelServiceImpl) QueryUserFavoriteCountList(ctx context.Context, req *api.FavoriteModelQueryUserFavoriteCountListRequest) (resp *api.FavoriteModelQueryUserFavoriteCountListResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryVideoFavoriteCountList implements the FavoriteModelServiceImpl interface.
func (s *FavoriteModelServiceImpl) QueryVideoFavoriteCountList(ctx context.Context, req *api.FavoriteModelQueryVideoFavoriteCountListRequest) (resp *api.FavoriteModelQueryVideoFavoriteCountListResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryVideoFavoriteCount implements the FavoriteModelServiceImpl interface.
func (s *FavoriteModelServiceImpl) QueryVideoFavoriteCount(ctx context.Context, req *api.FavoriteModelQueryVideoFavoriteCountRequest) (resp *api.FavoriteModelQueryVideoFavoriteCountResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryIsUserFavoriteList implements the FavoriteModelServiceImpl interface.
func (s *FavoriteModelServiceImpl) QueryIsUserFavoriteList(ctx context.Context, req *api.FavoriteModelQueryIsUserFavoriteListRequest) (resp *api.FavoriteModelQueryIsUserFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryIsUserFavorite implements the FavoriteModelServiceImpl interface.
func (s *FavoriteModelServiceImpl) QueryIsUserFavorite(ctx context.Context, req *api.FavoriteModelQueryIsUserFavoriteRequest) (resp *api.FavoriteModelQueryIsUserFavoriteResponse, err error) {
	// TODO: Your code here...
	return
}
