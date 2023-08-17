package core

import (
	"context"
	"time"
	api "videoModel/core/kitex_gen/api"
	"videoModel/model"
)

// VideoModelProtoBufImpl implements the last service interface defined in the IDL.
type VideoModelProtoBufImpl struct{}

// CreateVideo implements the VideoModelProtoBufImpl interface.
func (s *VideoModelProtoBufImpl) CreateVideo(ctx context.Context, req *api.VideoModelCreateVideoRequest) (resp *api.VideoModelCreateVideoResponse, err error) {
	// TODO: Your code here...
	//写入数据库
	t := time.Now()
	video := &model.Video{
		AuthorID:    req.AuthorId,
		PlayUrl:     req.PlayUrl,
		CoverUrl:    req.CoverUrl,
		Title:       req.Title,
		CreatedTime: &t,
		UpdatedTime: &t,
	}
	err = model.CreateVideo(video)

	if err != nil {
		errStr := "数据库创建video对象失败"
		resp = &api.VideoModelCreateVideoResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
		}
		return
	}

	//正常返回
	resp = &api.VideoModelCreateVideoResponse{
		StatusCode: 0,
	}
	return
}

// QueryAuthorWorkCount implements the VideoModelProtoBufImpl interface.
func (s *VideoModelProtoBufImpl) QueryAuthorWorkCount(ctx context.Context, req *api.VideoModelQueryAuthorWorkCountRequest) (resp *api.VideoModelQueryAuthorWorkCountResponse, err error) {
	// TODO: Your code here...
	count, err := model.QueryAuthorWorkCount(req.AuthorId)
	if err != nil {
		errStr := "查询作者作品数失败"
		resp = &api.VideoModelQueryAuthorWorkCountResponse{
			WorkCount:  0,
			StatusMsg:  &errStr,
			StatusCode: -1,
		}
		return
	}
	resp = &api.VideoModelQueryAuthorWorkCountResponse{
		WorkCount: uint32(count),
	}
	return
}

// QueryAuthorVideoList implements the VideoModelProtoBufImpl interface.
func (s *VideoModelProtoBufImpl) QueryAuthorVideoList(ctx context.Context, req *api.VideoModelQueryAuthorVideoListRequest) (resp *api.VideoModelQueryAuthorVideoListResponse, err error) {
	// TODO: Your code here...
	ormVideos, err := model.QueryAuthorVideoList(req.AuthorId)
	if err != nil {
		errStr := "数据库查询失败"
		resp = &api.VideoModelQueryAuthorVideoListResponse{
			VideoBaseInfo: nil,
			StatusMsg:     &errStr,
			StatusCode:    -1,
		}
		return
	}

	resp = &api.VideoModelQueryAuthorVideoListResponse{
		VideoBaseInfo: BuildApiBaseVideoInfoList(ormVideos),
	}
	return
}

func BuildApiBaseVideoInfoList(videos []*model.Video) []*api.VideoBaseInfo {
	var videoList []*api.VideoBaseInfo

	//遍历model中的video结构体数组，批量扩孔ApiVideoList
	for _, video := range videos {
		videoList = append(videoList, BuildApiBaseInfoVideo(video))
	}
	return videoList
}

func BuildApiBaseInfoVideo(video *model.Video) *api.VideoBaseInfo {
	//uid, _ := strconv.ParseInt(strconv.FormatUint(video.AuthorID, 10), 10, 64)
	//author := BuildApiUser(uid)
	return &api.VideoBaseInfo{
		VideoId:  video.VideoID,
		AuthorId: video.AuthorID,
		PlayUrl:  video.PlayUrl,
		CoverUrl: video.CoverUrl,
		Title:    video.Title,
	}
}

// QueryVideoList implements the VideoModelProtoBufImpl interface.
func (s *VideoModelProtoBufImpl) QueryVideoList(ctx context.Context, req *api.VideoModelQueryVideoListRequest) (resp *api.VideoModelQueryVideoListResponse, err error) {
	// TODO: Your code here...
	ormVideos, _ := model.GetVideosByIds(req.VideoIdList)

	resp = &api.VideoModelQueryVideoListResponse{
		VideoBaseInfoList: BuildApiBaseVideoInfoList(ormVideos),
	}
	return
}

// QueryVideo implements the VideoModelProtoBufImpl interface.
func (s *VideoModelProtoBufImpl) QueryVideo(ctx context.Context, req *api.VideoModelQueryVideoRequest) (resp *api.VideoModelQueryVideoResponse, err error) {
	// TODO: Your code here...
	ormVideo, _ := model.QuerySingleVideo(req.VideoId)

	resp = &api.VideoModelQueryVideoResponse{
		VideoBaseInfo: BuildApiBaseInfoVideo(ormVideo),
	}
	return
}

// QueryVideoFeed implements the VideoModelProtoBufImpl interface.
func (s *VideoModelProtoBufImpl) QueryVideoFeed(ctx context.Context, req *api.VideoModelQueryVideoFeedRequest) (resp *api.VideoModelQueryVideoFeedResponse, err error) {
	// TODO: Your code here...

	//转换并格式化时间
	t := time.Unix(req.NextTime/1000, 0)
	format := "2006-01-02 15:04:05"
	searchTime := t.Format(format)

	ormVideos, _ := model.QueryVideoFeedByLastTimeAndLimit(&searchTime, 5)
	resp = &api.VideoModelQueryVideoFeedResponse{
		VideoBaseInfoList: BuildApiBaseVideoInfoList(ormVideos),
	}
	return
}
