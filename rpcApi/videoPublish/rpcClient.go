package videoPublish

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"rpcApi/videoPublish/api"
	"rpcApi/videoPublish/api/videopublishservice"
)

var videoPublishRpcClient videopublishservice.Client

func InitVideoPublishRpcClient() videopublishservice.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	videoPublishRpcClient, err = videopublishservice.NewClient("videoPublish", client.WithResolver(r))

	if err != nil {
		log.Fatal("videoPublish微服务rpcClient初始化链接失败")
		log.Fatal(err)
		return nil
	}
	return videoPublishRpcClient
}

func PublishVideo(userID int64, data []byte, title string) error {
	rpcReq := &api.VideoPublishActionRequest{
		UserId: userID,
		Data:   data,
		Title:  title,
	}

	_, err := videoPublishRpcClient.PublishVideo(context.Background(), rpcReq)

	return err
}
