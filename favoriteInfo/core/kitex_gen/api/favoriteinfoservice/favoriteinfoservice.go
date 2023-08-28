// Code generated by Kitex v0.6.2. DO NOT EDIT.

package favoriteinfoservice

import (
	"context"
	api "favoriteInfo/core/kitex_gen/api"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteInfoServiceServiceInfo
}

var favoriteInfoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "favoriteInfoService"
	handlerType := (*api.FavoriteInfoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"queryFavoriteList": kitex.NewMethodInfo(queryFavoriteListHandler, newQueryFavoriteListArgs, newQueryFavoriteListResult, false),
		"favoriteAction":    kitex.NewMethodInfo(favoriteActionHandler, newFavoriteActionArgs, newFavoriteActionResult, false),
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

func queryFavoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.FavoriteInfoQueryFavoriteListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FavoriteInfoService).QueryFavoriteList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryFavoriteListArgs:
		success, err := handler.(api.FavoriteInfoService).QueryFavoriteList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryFavoriteListResult)
		realResult.Success = success
	}
	return nil
}
func newQueryFavoriteListArgs() interface{} {
	return &QueryFavoriteListArgs{}
}

func newQueryFavoriteListResult() interface{} {
	return &QueryFavoriteListResult{}
}

type QueryFavoriteListArgs struct {
	Req *api.FavoriteInfoQueryFavoriteListRequest
}

func (p *QueryFavoriteListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.FavoriteInfoQueryFavoriteListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryFavoriteListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryFavoriteListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryFavoriteListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryFavoriteListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryFavoriteListArgs) Unmarshal(in []byte) error {
	msg := new(api.FavoriteInfoQueryFavoriteListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryFavoriteListArgs_Req_DEFAULT *api.FavoriteInfoQueryFavoriteListRequest

func (p *QueryFavoriteListArgs) GetReq() *api.FavoriteInfoQueryFavoriteListRequest {
	if !p.IsSetReq() {
		return QueryFavoriteListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryFavoriteListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryFavoriteListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryFavoriteListResult struct {
	Success *api.FavoriteInfoQueryFavoriteListResponse
}

var QueryFavoriteListResult_Success_DEFAULT *api.FavoriteInfoQueryFavoriteListResponse

func (p *QueryFavoriteListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.FavoriteInfoQueryFavoriteListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryFavoriteListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryFavoriteListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryFavoriteListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryFavoriteListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryFavoriteListResult) Unmarshal(in []byte) error {
	msg := new(api.FavoriteInfoQueryFavoriteListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryFavoriteListResult) GetSuccess() *api.FavoriteInfoQueryFavoriteListResponse {
	if !p.IsSetSuccess() {
		return QueryFavoriteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryFavoriteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.FavoriteInfoQueryFavoriteListResponse)
}

func (p *QueryFavoriteListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryFavoriteListResult) GetResult() interface{} {
	return p.Success
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.FavoriteInfoFavoriteActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FavoriteInfoService).FavoriteAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteActionArgs:
		success, err := handler.(api.FavoriteInfoService).FavoriteAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteActionResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteActionArgs() interface{} {
	return &FavoriteActionArgs{}
}

func newFavoriteActionResult() interface{} {
	return &FavoriteActionResult{}
}

type FavoriteActionArgs struct {
	Req *api.FavoriteInfoFavoriteActionRequest
}

func (p *FavoriteActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.FavoriteInfoFavoriteActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoriteActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoriteActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoriteActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteActionArgs) Unmarshal(in []byte) error {
	msg := new(api.FavoriteInfoFavoriteActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteActionArgs_Req_DEFAULT *api.FavoriteInfoFavoriteActionRequest

func (p *FavoriteActionArgs) GetReq() *api.FavoriteInfoFavoriteActionRequest {
	if !p.IsSetReq() {
		return FavoriteActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteActionArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FavoriteActionArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FavoriteActionResult struct {
	Success *api.FavoriteInfoFavoriteActionResponse
}

var FavoriteActionResult_Success_DEFAULT *api.FavoriteInfoFavoriteActionResponse

func (p *FavoriteActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.FavoriteInfoFavoriteActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoriteActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoriteActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoriteActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteActionResult) Unmarshal(in []byte) error {
	msg := new(api.FavoriteInfoFavoriteActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteActionResult) GetSuccess() *api.FavoriteInfoFavoriteActionResponse {
	if !p.IsSetSuccess() {
		return FavoriteActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.FavoriteInfoFavoriteActionResponse)
}

func (p *FavoriteActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FavoriteActionResult) GetResult() interface{} {
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

func (p *kClient) QueryFavoriteList(ctx context.Context, Req *api.FavoriteInfoQueryFavoriteListRequest) (r *api.FavoriteInfoQueryFavoriteListResponse, err error) {
	var _args QueryFavoriteListArgs
	_args.Req = Req
	var _result QueryFavoriteListResult
	if err = p.c.Call(ctx, "queryFavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteAction(ctx context.Context, Req *api.FavoriteInfoFavoriteActionRequest) (r *api.FavoriteInfoFavoriteActionResponse, err error) {
	var _args FavoriteActionArgs
	_args.Req = Req
	var _result FavoriteActionResult
	if err = p.c.Call(ctx, "favoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}