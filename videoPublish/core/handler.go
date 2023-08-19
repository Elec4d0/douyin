package core

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	api "videoPublish/core/kitex_gen/api"
	model "videoPublish/rpcApi/videoModel"
	"videoPublish/tools/ffmepg"
	"videoPublish/tools/oss"
)

// VideoPublishServiceImpl implements the last service interface defined in the IDL.
type VideoPublishServiceImpl struct{}

// PublishVideo implements the VideoPublishServiceImpl interface.
func (s *VideoPublishServiceImpl) PublishVideo(ctx context.Context, req *api.VideoPublishActionRequest) (resp *api.VideoPublishActionResponse, err error) {
	// TODO: Your code here...
	oss.InitOss()
	fmt.Println(req.UserId)

	//给视频生成唯一UUID
	uuid := uuid.New().String()

	//上传视频
	videoFileName := uuid + ".mp4"
	oss.UploadVideo(req.Data, videoFileName)
	videoUrl := "http://douyin.g324.asia:9000/video/"
	ossVideoUrl := videoUrl + videoFileName

	//视频封面用ffmpeg截取第一帧后上传至OSS
	//通过视频URL获取封面
	jepgByte, err := ffmepg.GetVideoFirstFrameBytes(videoUrl)
	if err != nil {
		log.Println("获取封面失败")
	}

	//上传封面
	coverFileName := uuid + ".jpg"
	oss.UploadJepg(jepgByte, coverFileName)
	coverUrl := "http://douyin.g324.asia:9000/jepg/"
	ossCoverUrl := coverUrl + coverFileName

	//写入videoModel层
	model.CreateVideo(req.UserId, ossVideoUrl, ossCoverUrl, req.Title)
	StatusCode := "上传成功"
	resp = &api.VideoPublishActionResponse{
		StatusCode: 0,
		StatusMsg:  &StatusCode,
	}
	return
}
