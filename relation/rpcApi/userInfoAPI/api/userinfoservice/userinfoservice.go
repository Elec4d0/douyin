// Code generated by Kitex v0.6.2. DO NOT EDIT.

package userinfoservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	api 	"relation/rpcApi/userInfoAPI/api"
)

func serviceInfo() *kitex.ServiceInfo {
	return userInfoServiceServiceInfo
}

var userInfoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserInfoService"
	handlerType := (*api.UserInfoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetFullUserInfo":     kitex.NewMethodInfo(getFullUserInfoHandler, newGetFullUserInfoArgs, newGetFullUserInfoResult, false),
		"GetFullUserInfoList": kitex.NewMethodInfo(getFullUserInfoListHandler, newGetFullUserInfoListArgs, newGetFullUserInfoListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "services",
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

func getFullUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinUserGetFullUserInfoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.UserInfoService).GetFullUserInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetFullUserInfoArgs:
		success, err := handler.(api.UserInfoService).GetFullUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFullUserInfoResult)
		realResult.Success = success
	}
	return nil
}
func newGetFullUserInfoArgs() interface{} {
	return &GetFullUserInfoArgs{}
}

func newGetFullUserInfoResult() interface{} {
	return &GetFullUserInfoResult{}
}

type GetFullUserInfoArgs struct {
	Req *api.DouyinUserGetFullUserInfoRequest
}

func (p *GetFullUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinUserGetFullUserInfoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetFullUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetFullUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetFullUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFullUserInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetFullUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserGetFullUserInfoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFullUserInfoArgs_Req_DEFAULT *api.DouyinUserGetFullUserInfoRequest

func (p *GetFullUserInfoArgs) GetReq() *api.DouyinUserGetFullUserInfoRequest {
	if !p.IsSetReq() {
		return GetFullUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFullUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetFullUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetFullUserInfoResult struct {
	Success *api.DouyinUserGetFullUserInfoResponse
}

var GetFullUserInfoResult_Success_DEFAULT *api.DouyinUserGetFullUserInfoResponse

func (p *GetFullUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinUserGetFullUserInfoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetFullUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetFullUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetFullUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFullUserInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetFullUserInfoResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserGetFullUserInfoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFullUserInfoResult) GetSuccess() *api.DouyinUserGetFullUserInfoResponse {
	if !p.IsSetSuccess() {
		return GetFullUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFullUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinUserGetFullUserInfoResponse)
}

func (p *GetFullUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetFullUserInfoResult) GetResult() interface{} {
	return p.Success
}

func getFullUserInfoListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinUserGetFullUserInfoListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.UserInfoService).GetFullUserInfoList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetFullUserInfoListArgs:
		success, err := handler.(api.UserInfoService).GetFullUserInfoList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFullUserInfoListResult)
		realResult.Success = success
	}
	return nil
}
func newGetFullUserInfoListArgs() interface{} {
	return &GetFullUserInfoListArgs{}
}

func newGetFullUserInfoListResult() interface{} {
	return &GetFullUserInfoListResult{}
}

type GetFullUserInfoListArgs struct {
	Req *api.DouyinUserGetFullUserInfoListRequest
}

func (p *GetFullUserInfoListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinUserGetFullUserInfoListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetFullUserInfoListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetFullUserInfoListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetFullUserInfoListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFullUserInfoListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetFullUserInfoListArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserGetFullUserInfoListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFullUserInfoListArgs_Req_DEFAULT *api.DouyinUserGetFullUserInfoListRequest

func (p *GetFullUserInfoListArgs) GetReq() *api.DouyinUserGetFullUserInfoListRequest {
	if !p.IsSetReq() {
		return GetFullUserInfoListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFullUserInfoListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetFullUserInfoListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetFullUserInfoListResult struct {
	Success *api.DouyinUserGetFullUserInfoListResponse
}

var GetFullUserInfoListResult_Success_DEFAULT *api.DouyinUserGetFullUserInfoListResponse

func (p *GetFullUserInfoListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinUserGetFullUserInfoListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetFullUserInfoListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetFullUserInfoListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetFullUserInfoListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFullUserInfoListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetFullUserInfoListResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserGetFullUserInfoListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFullUserInfoListResult) GetSuccess() *api.DouyinUserGetFullUserInfoListResponse {
	if !p.IsSetSuccess() {
		return GetFullUserInfoListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFullUserInfoListResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinUserGetFullUserInfoListResponse)
}

func (p *GetFullUserInfoListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetFullUserInfoListResult) GetResult() interface{} {
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

func (p *kClient) GetFullUserInfo(ctx context.Context, Req *api.DouyinUserGetFullUserInfoRequest) (r *api.DouyinUserGetFullUserInfoResponse, err error) {
	var _args GetFullUserInfoArgs
	_args.Req = Req
	var _result GetFullUserInfoResult
	if err = p.c.Call(ctx, "GetFullUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFullUserInfoList(ctx context.Context, Req *api.DouyinUserGetFullUserInfoListRequest) (r *api.DouyinUserGetFullUserInfoListResponse, err error) {
	var _args GetFullUserInfoListArgs
	_args.Req = Req
	var _result GetFullUserInfoListResult
	if err = p.c.Call(ctx, "GetFullUserInfoList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
