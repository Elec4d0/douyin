package main

import (
	"comment/server/protos/kitex_gen/api"
	"comment/server/protos/kitex_gen/api/commentserver"
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func test3() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	cli, err := commentserver.NewClient("CommentServer", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	var ids []int64 = []int64{1, 2}
	req := &api.DouyinCommentserverCommentallcountRequest{
		VideoIds: ids,
	}
	resp, err := cli.CommentAllCount(context.Background(), req)
	/*jsonpbMarshaler := &jsonpb.Marshaler{
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

	fmt.Println("----end test-------")*/
	log.Println(resp)
}
