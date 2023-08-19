package main

import (
	"bytes"
	"comment/services/protos/kitex_gen/api"
	"comment/services/protos/kitex_gen/api/commentserver"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/gogo/protobuf/jsonpb"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func h() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	cli, err := commentserver.NewClient("CommentServer", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	commentID := int64(2)
	req := &api.DouyinCommentActionRequest{
		ActionType: 2,
		UserId:     1,
		VideoId:    1,
		CommentId:  &commentID,
	}
	resp, err := cli.CommentAction(context.Background(), req)
	jsonpbMarshaler := &jsonpb.Marshaler{
		EnumsAsInts:  true, //是否将枚举值设定为整数，而不是字符串类型
		EmitDefaults: true, //是否将字段值为空的渲染到JSON结构中
		OrigName:     true, //是否使用原生的proto协议中的字段
	}

	var buffer bytes.Buffer
	err = jsonpbMarshaler.Marshal(&buffer, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(buffer.Bytes()))

	fmt.Println("----end test-------")
}
