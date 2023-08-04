package protos

import (
	"context"
	"fmt"
	"time"
	"video/model"
	api "video/services/protos/kitex_gen/api"
	"video/tools/oss"

	"video/rpcClient"
)

// FeedProtoBufImpl implements the last service interface defined in the IDL.
type FeedProtoBufImpl struct{}

// GetFeed implements the FeedProtoBufImpl interface.
func (s *FeedProtoBufImpl) GetFeed(ctx context.Context, req *api.DouyinFeedRequest) (resp *api.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	//未登录用户返回视频流，无需关心视频是否点赞，默认为未登录
	var userId int64 = req.UserId

	//格式化时间，使其规范化用于mysql查询
	//int64 --> time --> str
	lastTime := *req.LatestTime
	var t time.Time

	if lastTime == -1 {
		t = time.Now()
	} else {
		t = time.Unix(lastTime/1000, 0)
	}
	t = time.Now()

	format := "2006-01-02 15:04:05"
	searchTime := t.Format(format)

	ormVideos, err := model.QueryVideoFeedByLastTimeAndLimit(&searchTime, 5)

	if err != nil {
		str := "Feed流查询失败"
		fmt.Println(str)
		resp.VideoList = nil
		resp.StatusCode = -1
		resp.StatusMsg = &str
		resp.NextTime = req.LatestTime
		return
	}

	StatusMsg := "查询成功"
	NextTime := time.Time.Unix(ormVideos[len(ormVideos)-1].CreatedTime)

	//gorm --> api兼容层转换
	apiVideoList := BuildApiVideoList(ormVideos)

	//构造作者ID列表
	authorIdList := make([]uint64, len(ormVideos))
	for i, ormVideo := range ormVideos {
		authorIdList[i] = ormVideo.AuthorID
	}

	//根据作者ID，通过User微服务更新Api视频列表的Author对象
	UpdateFeedViedosAuthor(apiVideoList, authorIdList)

	//未登录状态，不需要更新：1、是否点赞视频 2、是否关注作者
	if userId < 0 {
		//resp.VideoList = apiVideoList
	}

	//通过relation 微服务查询关注信息，并更新是否关注视频作者

	//通过favorite 微服务查询点赞信息，并更新是否点赞视频

	resp = &api.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  &StatusMsg,
		VideoList:  apiVideoList,
		NextTime:   &NextTime,
	}
	return
}

// PublishVideo implements the FeedProtoBufImpl interface.
func (s *FeedProtoBufImpl) PublishVideo(ctx context.Context, req *api.DouyinPublishActionRequest) (resp *api.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	oss.InitOss()
	fmt.Println(req.UserId)
	//用三方工具hash，给视频生成唯一UUID
	videoFileName := req.Title

	//上传文件后，拼接视频和视频封面Url
	//视频封面用ffmpeg截取第一帧后上传至OSS
	oss.UploadVideo(req.Data, videoFileName)
	baseUrl := "http://douyin.g324.asia:9000/video/"
	ossVideoUrl := baseUrl + videoFileName
	ossCoverUrl := ""

	//作者的用户ID需从Token微服务接口获取
	video := &model.Video{
		AuthorID:      uint64(req.UserId),
		PlayUrl:       ossVideoUrl,
		CoverUrl:      ossCoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
		CreatedTime:   time.Now(),
	}
	model.CreateVideo(video)

	StatusCode := "上传成功"
	resp = &api.DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg:  &StatusCode,
	}
	return
}

// GetAuthorVideoList implements the FeedProtoBufImpl interface.
func (s *FeedProtoBufImpl) GetAuthorVideoList(ctx context.Context, req *api.DouyinPublishListRequest) (resp *api.DouyinPublishListResponse, err error) {
	// TODO: Your code here...

	//校验UserId是否为负
	if req.UserId < 0 {
		str := "UserId为负值，非法"
		resp = &api.DouyinPublishListResponse{
			VideoList:  nil,
			StatusCode: -1,
			StatusMsg:  &str,
		}
		return
	}
	authorId := uint64(req.UserId)

	videos, err := model.QueryAuthorVideoList(authorId)

	if err != nil {
		str := "数据库查询出错，请联系管理员"
		resp = &api.DouyinPublishListResponse{
			VideoList:  nil,
			StatusCode: -2,
			StatusMsg:  &str,
		}
		return
	}

	//构造API层级的结构体，并用gorm层级的结构体遍历填充，胶水代码
	var apiVideoList []*api.Video
	apiVideoList = BuildApiVideoList(videos)
	//fmt.Println(apiVideoList)
	UpdateApiViedosAuthor(apiVideoList, authorId)
	//fmt.Println(apiVideoList)
	//查询全程正常，返回RPC查询结果
	str := "用户视频列表查询成功"
	sc := 0
	resp = &api.DouyinPublishListResponse{
		VideoList:  apiVideoList,
		StatusCode: int32(sc),
		StatusMsg:  &str,
	}

	fmt.Println(resp)
	return
}

func BuildApiVideoList(videos []*model.Video) []*api.Video {
	var videoList []*api.Video

	//遍历model中的video结构体数组，批量扩孔ApiVideoList
	for _, video := range videos {
		videoList = append(videoList, BuildApiVideo(video))
	}
	return videoList
}

func BuildApiVideo(video *model.Video) *api.Video {
	//uid, _ := strconv.ParseInt(strconv.FormatUint(video.AuthorID, 10), 10, 64)
	//author := BuildApiUser(uid)
	return &api.Video{
		Id:            int64(video.VideoID),
		Author:        nil,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: int64(video.FavoriteCount),
		CommentCount:  int64(video.CommentCount),
		IsFavorite:    false,
		Title:         video.Title,
	}
}

/*
func BuildApiUser(UserId int64) *api.User {

	return &api.User{
		Id:              UserId,
		Name:            "",
		FollowCount:     nil,
		FollowerCount:   nil,
		IsFollow:        false,
		Avatar:          nil,
		BackgroundImage: nil,
		Signature:       nil,
		TotalFavorited:  nil,
		WorkCount:       nil,
		FavoriteCount:   nil,
	}
}*/

func UpdateFeedViedosAuthor(videolist []*api.Video, authorIdList []uint64) {
	for i, video := range videolist {
		authorId := authorIdList[i]
		//fmt.Println(authorId)
		usr, err := rpcClient.GetUserInFo(authorId)
		if err != nil {
			continue
		}
		video.Author = usr
	}
}

func UpdateApiViedosAuthor(apiVideoList []*api.Video, authorId uint64) {
	usr, err := rpcClient.GetUserInFo(authorId)
	if err != nil {
		fmt.Println("在publish/list通过作者id获取作者实例出错")
		return
	}
	for _, video := range apiVideoList {
		video.Author = usr
	}
}
