package core

import (
	"context"
	"log"
	"time"
	api "videoModel/core/kitex_gen/api"
	"videoModel/model"
)

// VideoModelServiceImpl implements the last service interface defined in the IDL.
type VideoModelServiceImpl struct{}

// CreateVideo implements the VideoModelServiceImpl interface.
func (s *VideoModelServiceImpl) CreateVideo(ctx context.Context, req *api.VideoModelCreateVideoRequest) (resp *api.VideoModelCreateVideoResponse, err error) {
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

// QueryAuthorWorkCount implements the VideoModelServiceImpl interface.
func (s *VideoModelServiceImpl) QueryAuthorWorkCount(ctx context.Context, req *api.VideoModelQueryAuthorWorkCountRequest) (resp *api.VideoModelQueryAuthorWorkCountResponse, err error) {
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
		WorkCount: count,
	}
	return
}

// QueryAuthorVideoIDList implements the VideoModelServiceImpl interface.
func (s *VideoModelServiceImpl) QueryAuthorVideoIDList(ctx context.Context, req *api.VideoModelQueryAuthorVideoIdListRequest) (resp *api.VideoModelQueryAuthorVideoIdListResponse, err error) {
	// TODO: Your code here...
	ormVideos, err := model.QueryAuthorVideoList(req.AuthorId)
	if err != nil {
		errStr := "数据库查询失败"
		resp = &api.VideoModelQueryAuthorVideoIdListResponse{
			VideoIdList: nil,
			StatusMsg:   &errStr,
			StatusCode:  -1,
		}
		return
	}

	var videoIDList = make([]int64, len(ormVideos))

	for i, ormVideo := range ormVideos {
		videoIDList[i] = ormVideo.VideoID
	}
	resp = &api.VideoModelQueryAuthorVideoIdListResponse{
		VideoIdList: videoIDList,
	}
	return
}

func BuildApiBaseVideoInfoList(ormVideos []*model.Video) []*api.VideoModel {
	n := len(ormVideos)
	apiVideos := make([]*api.VideoModel, n, n)

	//遍历model中的video结构体数组，批量扩孔ApiVideoList
	for i, video := range ormVideos {
		//fmt.Println(video)
		apiVideos[i] = BuildApiBaseInfoVideo(video)
	}
	//fmt.Println(apiVideos)
	return apiVideos
}

func BuildApiBaseInfoVideo(video *model.Video) *api.VideoModel {
	//uid, _ := strconv.ParseInt(strconv.FormatUint(video.AuthorID, 10), 10, 64)
	//author := BuildApiUser(uid)
	if video == nil {
		return nil
	}
	return &api.VideoModel{
		VideoId:  video.VideoID,
		AuthorId: video.AuthorID,
		PlayUrl:  video.PlayUrl,
		CoverUrl: video.CoverUrl,
		Title:    video.Title,
	}
}

// QueryVideoList implements the VideoModelServiceImpl interface.
func (s *VideoModelServiceImpl) QueryVideoList(ctx context.Context, req *api.VideoModelQueryVideoListRequest) (resp *api.VideoModelQueryVideoListResponse, err error) {
	// TODO: Your code here...
	ormVideos, _ := model.GetVideosByIds(req.VideoIdList)

	resp = &api.VideoModelQueryVideoListResponse{
		VideoModelList: BuildApiBaseVideoInfoList(ormVideos),
	}
	return
}

// QueryVideo implements the VideoModelServiceImpl interface.
func (s *VideoModelServiceImpl) QueryVideo(ctx context.Context, req *api.VideoModelQueryVideoRequest) (resp *api.VideoModelQueryVideoResponse, err error) {
	// TODO: Your code here...
	ormVideo, _ := model.QuerySingleVideo(req.VideoId)

	resp = &api.VideoModelQueryVideoResponse{
		VideoModel: BuildApiBaseInfoVideo(ormVideo),
	}
	return
}

// QueryVideoFeed implements the VideoModelServiceImpl interface.
func (s *VideoModelServiceImpl) QueryVideoFeed(ctx context.Context, req *api.VideoModelQueryVideoFeedRequest) (resp *api.VideoModelQueryVideoFeedResponse, err error) {
	// TODO: Your code here...
	log.Println("接收到时间戳:", req.NextTime)
	//转换并格式化时间
	t := time.Unix(req.NextTime/1000, 0)
	format := "2006-01-02 15:04:05"
	searchTime := t.Format(format)
	log.Println("Feed流格式化时间：", searchTime)
	ormVideoList, _ := model.QueryVideoFeedByLastTimeAndLimit(&searchTime, req.Limit)
	log.Println(ormVideoList)
	var videoIDList = make([]int64, len(ormVideoList))
	var createTimeList = make([]int64, len(ormVideoList))
	for i, ormVideo := range ormVideoList {
		videoIDList[i] = ormVideo.VideoID
		createTimeList[i] = ormVideo.CreatedTime.Unix() * 1000
	}
	resp = &api.VideoModelQueryVideoFeedResponse{
		VideoIdList:    videoIDList,
		CreateTimeList: createTimeList,
	}
	return
}

// QueryAuthorWorkCountList implements the VideoModelServiceImpl interface.
func (s *VideoModelServiceImpl) QueryAuthorWorkCountList(ctx context.Context, req *api.VideoModelQueryAuthorWorkCountListRequest) (resp *api.VideoModelQueryAuthorWorkCountListResponse, err error) {
	// TODO: Your code here...
	authorIDlist := req.AuthorIdList

	n := len(authorIDlist)
	var workCountList = make([]int32, n, n)

	for i, authorID := range authorIDlist {
		count, _ := model.QueryAuthorWorkCount(authorID)
		workCountList[i] = count
	}

	resp = &api.VideoModelQueryAuthorWorkCountListResponse{
		WorkCountList: workCountList,
	}
	return
}
