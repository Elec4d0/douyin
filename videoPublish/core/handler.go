package core

import (
	"context"
	"fmt"
	api "videoPublish/core/kitex_gen/api"
	model "videoPublish/rpcApi/videoModel"
	"videoPublish/tools/oss"
)

// VideoPublishServiceImpl implements the last service interface defined in the IDL.
type VideoPublishServiceImpl struct{}

// PublishVideo implements the VideoPublishServiceImpl interface.
func (s *VideoPublishServiceImpl) PublishVideo(ctx context.Context, req *api.VideoPublishActionRequest) (resp *api.VideoPublishActionResponse, err error) {
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
	model.CreateVideo(req.UserId, ossVideoUrl, ossCoverUrl, req.Title)
	StatusCode := "上传成功"
	resp = &api.VideoPublishActionResponse{
		StatusCode: 0,
		StatusMsg:  &StatusCode,
	}
	return
}
