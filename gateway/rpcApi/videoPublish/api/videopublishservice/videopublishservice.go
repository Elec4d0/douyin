// Code generated by Kitex v0.6.2. DO NOT EDIT.

package videopublishservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	api "gateway/rpcApi/videoPublish/api"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoPublishServiceServiceInfo
}

var videoPublishServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "videoPublishService"
	handlerType := (*api.VideoPublishService)(nil)
	methods := map[string]kitex.MethodInfo{
		"publishVideo": kitex.NewMethodInfo(publishVideoHandler, newPublishVideoArgs, newPublishVideoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "core",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func publishVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.VideoPublishActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.VideoPublishService).PublishVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishVideoArgs:
		success, err := handler.(api.VideoPublishService).PublishVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishVideoResult)
		realResult.Success = success
	}
	return nil
}
func newPublishVideoArgs() interface{} {
	return &PublishVideoArgs{}
}

func newPublishVideoResult() interface{} {
	return &PublishVideoResult{}
}

type PublishVideoArgs struct {
	Req *api.VideoPublishActionRequest
}

func (p *PublishVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.VideoPublishActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PublishVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PublishVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PublishVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishVideoArgs) Unmarshal(in []byte) error {
	msg := new(api.VideoPublishActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishVideoArgs_Req_DEFAULT *api.VideoPublishActionRequest

func (p *PublishVideoArgs) GetReq() *api.VideoPublishActionRequest {
	if !p.IsSetReq() {
		return PublishVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PublishVideoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PublishVideoResult struct {
	Success *api.VideoPublishActionResponse
}

var PublishVideoResult_Success_DEFAULT *api.VideoPublishActionResponse

func (p *PublishVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.VideoPublishActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PublishVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PublishVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PublishVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishVideoResult) Unmarshal(in []byte) error {
	msg := new(api.VideoPublishActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishVideoResult) GetSuccess() *api.VideoPublishActionResponse {
	if !p.IsSetSuccess() {
		return PublishVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.VideoPublishActionResponse)
}

func (p *PublishVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PublishVideoResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PublishVideo(ctx context.Context, Req *api.VideoPublishActionRequest) (r *api.VideoPublishActionResponse, err error) {
	var _args PublishVideoArgs
	_args.Req = Req
	var _result PublishVideoResult
	if err = p.c.Call(ctx, "publishVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
