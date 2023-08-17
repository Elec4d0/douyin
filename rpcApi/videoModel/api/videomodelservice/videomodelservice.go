// Code generated by Kitex v0.6.2. DO NOT EDIT.

package videomodelservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	api "rpcApi/videoModel/api"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoModelServiceServiceInfo
}

var videoModelServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "videoModelService"
	handlerType := (*api.VideoModelService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateVideo":          kitex.NewMethodInfo(createVideoHandler, newCreateVideoArgs, newCreateVideoResult, false),
		"QueryAuthorWorkCount": kitex.NewMethodInfo(queryAuthorWorkCountHandler, newQueryAuthorWorkCountArgs, newQueryAuthorWorkCountResult, false),
		"QueryAuthorVideoList": kitex.NewMethodInfo(queryAuthorVideoListHandler, newQueryAuthorVideoListArgs, newQueryAuthorVideoListResult, false),
		"QueryVideoList":       kitex.NewMethodInfo(queryVideoListHandler, newQueryVideoListArgs, newQueryVideoListResult, false),
		"QueryVideo":           kitex.NewMethodInfo(queryVideoHandler, newQueryVideoArgs, newQueryVideoResult, false),
		"QueryVideoFeed":       kitex.NewMethodInfo(queryVideoFeedHandler, newQueryVideoFeedArgs, newQueryVideoFeedResult, false),
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

func createVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.VideoModelCreateVideoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.VideoModelService).CreateVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateVideoArgs:
		success, err := handler.(api.VideoModelService).CreateVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateVideoResult)
		realResult.Success = success
	}
	return nil
}
func newCreateVideoArgs() interface{} {
	return &CreateVideoArgs{}
}

func newCreateVideoResult() interface{} {
	return &CreateVideoResult{}
}

type CreateVideoArgs struct {
	Req *api.VideoModelCreateVideoRequest
}

func (p *CreateVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.VideoModelCreateVideoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateVideoArgs) Unmarshal(in []byte) error {
	msg := new(api.VideoModelCreateVideoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateVideoArgs_Req_DEFAULT *api.VideoModelCreateVideoRequest

func (p *CreateVideoArgs) GetReq() *api.VideoModelCreateVideoRequest {
	if !p.IsSetReq() {
		return CreateVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateVideoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateVideoResult struct {
	Success *api.VideoModelCreateVideoResponse
}

var CreateVideoResult_Success_DEFAULT *api.VideoModelCreateVideoResponse

func (p *CreateVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.VideoModelCreateVideoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateVideoResult) Unmarshal(in []byte) error {
	msg := new(api.VideoModelCreateVideoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateVideoResult) GetSuccess() *api.VideoModelCreateVideoResponse {
	if !p.IsSetSuccess() {
		return CreateVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.VideoModelCreateVideoResponse)
}

func (p *CreateVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateVideoResult) GetResult() interface{} {
	return p.Success
}

func queryAuthorWorkCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.VideoModelQueryAuthorWorkCountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.VideoModelService).QueryAuthorWorkCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryAuthorWorkCountArgs:
		success, err := handler.(api.VideoModelService).QueryAuthorWorkCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryAuthorWorkCountResult)
		realResult.Success = success
	}
	return nil
}
func newQueryAuthorWorkCountArgs() interface{} {
	return &QueryAuthorWorkCountArgs{}
}

func newQueryAuthorWorkCountResult() interface{} {
	return &QueryAuthorWorkCountResult{}
}

type QueryAuthorWorkCountArgs struct {
	Req *api.VideoModelQueryAuthorWorkCountRequest
}

func (p *QueryAuthorWorkCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.VideoModelQueryAuthorWorkCountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryAuthorWorkCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryAuthorWorkCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryAuthorWorkCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryAuthorWorkCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryAuthorWorkCountArgs) Unmarshal(in []byte) error {
	msg := new(api.VideoModelQueryAuthorWorkCountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryAuthorWorkCountArgs_Req_DEFAULT *api.VideoModelQueryAuthorWorkCountRequest

func (p *QueryAuthorWorkCountArgs) GetReq() *api.VideoModelQueryAuthorWorkCountRequest {
	if !p.IsSetReq() {
		return QueryAuthorWorkCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryAuthorWorkCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryAuthorWorkCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryAuthorWorkCountResult struct {
	Success *api.VideoModelQueryAuthorWorkCountResponse
}

var QueryAuthorWorkCountResult_Success_DEFAULT *api.VideoModelQueryAuthorWorkCountResponse

func (p *QueryAuthorWorkCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.VideoModelQueryAuthorWorkCountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryAuthorWorkCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryAuthorWorkCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryAuthorWorkCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryAuthorWorkCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryAuthorWorkCountResult) Unmarshal(in []byte) error {
	msg := new(api.VideoModelQueryAuthorWorkCountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryAuthorWorkCountResult) GetSuccess() *api.VideoModelQueryAuthorWorkCountResponse {
	if !p.IsSetSuccess() {
		return QueryAuthorWorkCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryAuthorWorkCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.VideoModelQueryAuthorWorkCountResponse)
}

func (p *QueryAuthorWorkCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryAuthorWorkCountResult) GetResult() interface{} {
	return p.Success
}

func queryAuthorVideoListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.VideoModelQueryAuthorVideoListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.VideoModelService).QueryAuthorVideoList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryAuthorVideoListArgs:
		success, err := handler.(api.VideoModelService).QueryAuthorVideoList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryAuthorVideoListResult)
		realResult.Success = success
	}
	return nil
}
func newQueryAuthorVideoListArgs() interface{} {
	return &QueryAuthorVideoListArgs{}
}

func newQueryAuthorVideoListResult() interface{} {
	return &QueryAuthorVideoListResult{}
}

type QueryAuthorVideoListArgs struct {
	Req *api.VideoModelQueryAuthorVideoListRequest
}

func (p *QueryAuthorVideoListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.VideoModelQueryAuthorVideoListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryAuthorVideoListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryAuthorVideoListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryAuthorVideoListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryAuthorVideoListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryAuthorVideoListArgs) Unmarshal(in []byte) error {
	msg := new(api.VideoModelQueryAuthorVideoListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryAuthorVideoListArgs_Req_DEFAULT *api.VideoModelQueryAuthorVideoListRequest

func (p *QueryAuthorVideoListArgs) GetReq() *api.VideoModelQueryAuthorVideoListRequest {
	if !p.IsSetReq() {
		return QueryAuthorVideoListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryAuthorVideoListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryAuthorVideoListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryAuthorVideoListResult struct {
	Success *api.VideoModelQueryAuthorVideoListResponse
}

var QueryAuthorVideoListResult_Success_DEFAULT *api.VideoModelQueryAuthorVideoListResponse

func (p *QueryAuthorVideoListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.VideoModelQueryAuthorVideoListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryAuthorVideoListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryAuthorVideoListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryAuthorVideoListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryAuthorVideoListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryAuthorVideoListResult) Unmarshal(in []byte) error {
	msg := new(api.VideoModelQueryAuthorVideoListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryAuthorVideoListResult) GetSuccess() *api.VideoModelQueryAuthorVideoListResponse {
	if !p.IsSetSuccess() {
		return QueryAuthorVideoListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryAuthorVideoListResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.VideoModelQueryAuthorVideoListResponse)
}

func (p *QueryAuthorVideoListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryAuthorVideoListResult) GetResult() interface{} {
	return p.Success
}

func queryVideoListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.VideoModelQueryVideoListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.VideoModelService).QueryVideoList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryVideoListArgs:
		success, err := handler.(api.VideoModelService).QueryVideoList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryVideoListResult)
		realResult.Success = success
	}
	return nil
}
func newQueryVideoListArgs() interface{} {
	return &QueryVideoListArgs{}
}

func newQueryVideoListResult() interface{} {
	return &QueryVideoListResult{}
}

type QueryVideoListArgs struct {
	Req *api.VideoModelQueryVideoListRequest
}

func (p *QueryVideoListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.VideoModelQueryVideoListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryVideoListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryVideoListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryVideoListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryVideoListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryVideoListArgs) Unmarshal(in []byte) error {
	msg := new(api.VideoModelQueryVideoListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryVideoListArgs_Req_DEFAULT *api.VideoModelQueryVideoListRequest

func (p *QueryVideoListArgs) GetReq() *api.VideoModelQueryVideoListRequest {
	if !p.IsSetReq() {
		return QueryVideoListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryVideoListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryVideoListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryVideoListResult struct {
	Success *api.VideoModelQueryVideoListResponse
}

var QueryVideoListResult_Success_DEFAULT *api.VideoModelQueryVideoListResponse

func (p *QueryVideoListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.VideoModelQueryVideoListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryVideoListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryVideoListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryVideoListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryVideoListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryVideoListResult) Unmarshal(in []byte) error {
	msg := new(api.VideoModelQueryVideoListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryVideoListResult) GetSuccess() *api.VideoModelQueryVideoListResponse {
	if !p.IsSetSuccess() {
		return QueryVideoListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryVideoListResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.VideoModelQueryVideoListResponse)
}

func (p *QueryVideoListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryVideoListResult) GetResult() interface{} {
	return p.Success
}

func queryVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.VideoModelQueryVideoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.VideoModelService).QueryVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryVideoArgs:
		success, err := handler.(api.VideoModelService).QueryVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryVideoResult)
		realResult.Success = success
	}
	return nil
}
func newQueryVideoArgs() interface{} {
	return &QueryVideoArgs{}
}

func newQueryVideoResult() interface{} {
	return &QueryVideoResult{}
}

type QueryVideoArgs struct {
	Req *api.VideoModelQueryVideoRequest
}

func (p *QueryVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.VideoModelQueryVideoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryVideoArgs) Unmarshal(in []byte) error {
	msg := new(api.VideoModelQueryVideoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryVideoArgs_Req_DEFAULT *api.VideoModelQueryVideoRequest

func (p *QueryVideoArgs) GetReq() *api.VideoModelQueryVideoRequest {
	if !p.IsSetReq() {
		return QueryVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryVideoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryVideoResult struct {
	Success *api.VideoModelQueryVideoResponse
}

var QueryVideoResult_Success_DEFAULT *api.VideoModelQueryVideoResponse

func (p *QueryVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.VideoModelQueryVideoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryVideoResult) Unmarshal(in []byte) error {
	msg := new(api.VideoModelQueryVideoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryVideoResult) GetSuccess() *api.VideoModelQueryVideoResponse {
	if !p.IsSetSuccess() {
		return QueryVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.VideoModelQueryVideoResponse)
}

func (p *QueryVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryVideoResult) GetResult() interface{} {
	return p.Success
}

func queryVideoFeedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.VideoModelQueryVideoFeedRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.VideoModelService).QueryVideoFeed(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryVideoFeedArgs:
		success, err := handler.(api.VideoModelService).QueryVideoFeed(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryVideoFeedResult)
		realResult.Success = success
	}
	return nil
}
func newQueryVideoFeedArgs() interface{} {
	return &QueryVideoFeedArgs{}
}

func newQueryVideoFeedResult() interface{} {
	return &QueryVideoFeedResult{}
}

type QueryVideoFeedArgs struct {
	Req *api.VideoModelQueryVideoFeedRequest
}

func (p *QueryVideoFeedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.VideoModelQueryVideoFeedRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryVideoFeedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryVideoFeedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryVideoFeedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryVideoFeedArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryVideoFeedArgs) Unmarshal(in []byte) error {
	msg := new(api.VideoModelQueryVideoFeedRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryVideoFeedArgs_Req_DEFAULT *api.VideoModelQueryVideoFeedRequest

func (p *QueryVideoFeedArgs) GetReq() *api.VideoModelQueryVideoFeedRequest {
	if !p.IsSetReq() {
		return QueryVideoFeedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryVideoFeedArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryVideoFeedArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryVideoFeedResult struct {
	Success *api.VideoModelQueryVideoFeedResponse
}

var QueryVideoFeedResult_Success_DEFAULT *api.VideoModelQueryVideoFeedResponse

func (p *QueryVideoFeedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.VideoModelQueryVideoFeedResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryVideoFeedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryVideoFeedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryVideoFeedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryVideoFeedResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryVideoFeedResult) Unmarshal(in []byte) error {
	msg := new(api.VideoModelQueryVideoFeedResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryVideoFeedResult) GetSuccess() *api.VideoModelQueryVideoFeedResponse {
	if !p.IsSetSuccess() {
		return QueryVideoFeedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryVideoFeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.VideoModelQueryVideoFeedResponse)
}

func (p *QueryVideoFeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryVideoFeedResult) GetResult() interface{} {
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

func (p *kClient) CreateVideo(ctx context.Context, Req *api.VideoModelCreateVideoRequest) (r *api.VideoModelCreateVideoResponse, err error) {
	var _args CreateVideoArgs
	_args.Req = Req
	var _result CreateVideoResult
	if err = p.c.Call(ctx, "CreateVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryAuthorWorkCount(ctx context.Context, Req *api.VideoModelQueryAuthorWorkCountRequest) (r *api.VideoModelQueryAuthorWorkCountResponse, err error) {
	var _args QueryAuthorWorkCountArgs
	_args.Req = Req
	var _result QueryAuthorWorkCountResult
	if err = p.c.Call(ctx, "QueryAuthorWorkCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryAuthorVideoList(ctx context.Context, Req *api.VideoModelQueryAuthorVideoListRequest) (r *api.VideoModelQueryAuthorVideoListResponse, err error) {
	var _args QueryAuthorVideoListArgs
	_args.Req = Req
	var _result QueryAuthorVideoListResult
	if err = p.c.Call(ctx, "QueryAuthorVideoList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryVideoList(ctx context.Context, Req *api.VideoModelQueryVideoListRequest) (r *api.VideoModelQueryVideoListResponse, err error) {
	var _args QueryVideoListArgs
	_args.Req = Req
	var _result QueryVideoListResult
	if err = p.c.Call(ctx, "QueryVideoList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryVideo(ctx context.Context, Req *api.VideoModelQueryVideoRequest) (r *api.VideoModelQueryVideoResponse, err error) {
	var _args QueryVideoArgs
	_args.Req = Req
	var _result QueryVideoResult
	if err = p.c.Call(ctx, "QueryVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryVideoFeed(ctx context.Context, Req *api.VideoModelQueryVideoFeedRequest) (r *api.VideoModelQueryVideoFeedResponse, err error) {
	var _args QueryVideoFeedArgs
	_args.Req = Req
	var _result QueryVideoFeedResult
	if err = p.c.Call(ctx, "QueryVideoFeed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
