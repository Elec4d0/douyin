package messageAPI

import (
	"context"
	"fmt"
	"gateway/rpcApi/messageAPI/api"
	"gateway/rpcApi/messageAPI/api/messageservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var messageRpcClient messageservice.Client

func InitMessageRpcClient() messageservice.Client {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	messageRpcClient, err = messageservice.NewClient("MessageService", client.WithResolver(r))

	if err != nil {
		log.Println("网关层message 微服务初始化链接失败")
		log.Fatal(err)
		return nil
	}
	fmt.Println("message 微服务：初始化链接成功")
	return messageRpcClient
}

func MessageChat(fromUserId int64, toUserId int64, preMsgTime int64) (*api.DouyinMessageChatResponse, error) {
	rpcReq := &api.DouyinMessageChatRequest{
		FromUserId: fromUserId,
		ToUserId:   toUserId,
		PreMsgTime: preMsgTime,
	}
	fmt.Println(rpcReq)
	resp, err := messageRpcClient.MessageChat(context.Background(), rpcReq)

	if err != nil {
		log.Println(resp)
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

func RelationAction(fromUserId int64, toUserId int64, actionType int32, content string) (*api.DouyinRelationActionResponse, error) {
	rpcReq := &api.DouyinRelationActionRequest{
		FromUserId: fromUserId,
		ToUserId:   toUserId,
		ActionType: actionType,
		Content:    content,
	}
	fmt.Println(rpcReq)
	resp, err := messageRpcClient.RelationAction(context.Background(), rpcReq)

	if err != nil {
		log.Println(resp)
		log.Println(err)
		return nil, err
	}
	return resp, nil
}
