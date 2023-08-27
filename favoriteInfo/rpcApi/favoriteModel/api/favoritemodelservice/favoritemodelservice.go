// Code generated by Kitex v0.6.2. DO NOT EDIT.

package favoritemodelservice

import (
	"context"
	api "favoriteInfo/rpcApi/favoriteModel/api"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteModelServiceServiceInfo
}

var favoriteModelServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "favoriteModelService"
	handlerType := (*api.FavoriteModelService)(nil)
	methods := map[string]kitex.MethodInfo{
		"queryFavoriteList":           kitex.NewMethodInfo(queryFavoriteListHandler, newQueryFavoriteListArgs, newQueryFavoriteListResult, false),
		"favoriteAction":              kitex.NewMethodInfo(favoriteActionHandler, newFavoriteActionArgs, newFavoriteActionResult, false),
		"queryUserFavoriteCount":      kitex.NewMethodInfo(queryUserFavoriteCountHandler, newQueryUserFavoriteCountArgs, newQueryUserFavoriteCountResult, false),
		"queryUserFavoriteCountList":  kitex.NewMethodInfo(queryUserFavoriteCountListHandler, newQueryUserFavoriteCountListArgs, newQueryUserFavoriteCountListResult, false),
		"queryVideoFavoriteCountList": kitex.NewMethodInfo(queryVideoFavoriteCountListHandler, newQueryVideoFavoriteCountListArgs, newQueryVideoFavoriteCountListResult, false),
		"queryVideoFavoriteCount":     kitex.NewMethodInfo(queryVideoFavoriteCountHandler, newQueryVideoFavoriteCountArgs, newQueryVideoFavoriteCountResult, false),
		"queryIsUserFavoriteList":     kitex.NewMethodInfo(queryIsUserFavoriteListHandler, newQueryIsUserFavoriteListArgs, newQueryIsUserFavoriteListResult, false),
		"queryIsUserFavorite":         kitex.NewMethodInfo(queryIsUserFavoriteHandler, newQueryIsUserFavoriteArgs, newQueryIsUserFavoriteResult, false),
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
		req := new(api.FavoriteModelQueryFavoriteListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FavoriteModelService).QueryFavoriteList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryFavoriteListArgs:
		success, err := handler.(api.FavoriteModelService).QueryFavoriteList(ctx, s.Req)
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
	Req *api.FavoriteModelQueryFavoriteListRequest
}

func (p *QueryFavoriteListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.FavoriteModelQueryFavoriteListRequest)
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
	msg := new(api.FavoriteModelQueryFavoriteListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryFavoriteListArgs_Req_DEFAULT *api.FavoriteModelQueryFavoriteListRequest

func (p *QueryFavoriteListArgs) GetReq() *api.FavoriteModelQueryFavoriteListRequest {
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
	Success *api.FavoriteModelQueryFavoriteListResponse
}

var QueryFavoriteListResult_Success_DEFAULT *api.FavoriteModelQueryFavoriteListResponse

func (p *QueryFavoriteListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.FavoriteModelQueryFavoriteListResponse)
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
	msg := new(api.FavoriteModelQueryFavoriteListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryFavoriteListResult) GetSuccess() *api.FavoriteModelQueryFavoriteListResponse {
	if !p.IsSetSuccess() {
		return QueryFavoriteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryFavoriteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.FavoriteModelQueryFavoriteListResponse)
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
		req := new(api.FavoriteModelFavoriteActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FavoriteModelService).FavoriteAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteActionArgs:
		success, err := handler.(api.FavoriteModelService).FavoriteAction(ctx, s.Req)
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
	Req *api.FavoriteModelFavoriteActionRequest
}

func (p *FavoriteActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.FavoriteModelFavoriteActionRequest)
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
	msg := new(api.FavoriteModelFavoriteActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteActionArgs_Req_DEFAULT *api.FavoriteModelFavoriteActionRequest

func (p *FavoriteActionArgs) GetReq() *api.FavoriteModelFavoriteActionRequest {
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
	Success *api.FavoriteModelFavoriteActionResponse
}

var FavoriteActionResult_Success_DEFAULT *api.FavoriteModelFavoriteActionResponse

func (p *FavoriteActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.FavoriteModelFavoriteActionResponse)
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
	msg := new(api.FavoriteModelFavoriteActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteActionResult) GetSuccess() *api.FavoriteModelFavoriteActionResponse {
	if !p.IsSetSuccess() {
		return FavoriteActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.FavoriteModelFavoriteActionResponse)
}

func (p *FavoriteActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FavoriteActionResult) GetResult() interface{} {
	return p.Success
}

func queryUserFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.FavoriteModelQueryUserFavoriteCountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FavoriteModelService).QueryUserFavoriteCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryUserFavoriteCountArgs:
		success, err := handler.(api.FavoriteModelService).QueryUserFavoriteCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryUserFavoriteCountResult)
		realResult.Success = success
	}
	return nil
}
func newQueryUserFavoriteCountArgs() interface{} {
	return &QueryUserFavoriteCountArgs{}
}

func newQueryUserFavoriteCountResult() interface{} {
	return &QueryUserFavoriteCountResult{}
}

type QueryUserFavoriteCountArgs struct {
	Req *api.FavoriteModelQueryUserFavoriteCountRequest
}

func (p *QueryUserFavoriteCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.FavoriteModelQueryUserFavoriteCountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryUserFavoriteCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryUserFavoriteCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryUserFavoriteCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryUserFavoriteCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryUserFavoriteCountArgs) Unmarshal(in []byte) error {
	msg := new(api.FavoriteModelQueryUserFavoriteCountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryUserFavoriteCountArgs_Req_DEFAULT *api.FavoriteModelQueryUserFavoriteCountRequest

func (p *QueryUserFavoriteCountArgs) GetReq() *api.FavoriteModelQueryUserFavoriteCountRequest {
	if !p.IsSetReq() {
		return QueryUserFavoriteCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryUserFavoriteCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryUserFavoriteCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryUserFavoriteCountResult struct {
	Success *api.FavoriteModelQueryUserFavoriteCountResponse
}

var QueryUserFavoriteCountResult_Success_DEFAULT *api.FavoriteModelQueryUserFavoriteCountResponse

func (p *QueryUserFavoriteCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.FavoriteModelQueryUserFavoriteCountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryUserFavoriteCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryUserFavoriteCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryUserFavoriteCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryUserFavoriteCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryUserFavoriteCountResult) Unmarshal(in []byte) error {
	msg := new(api.FavoriteModelQueryUserFavoriteCountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryUserFavoriteCountResult) GetSuccess() *api.FavoriteModelQueryUserFavoriteCountResponse {
	if !p.IsSetSuccess() {
		return QueryUserFavoriteCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryUserFavoriteCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.FavoriteModelQueryUserFavoriteCountResponse)
}

func (p *QueryUserFavoriteCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryUserFavoriteCountResult) GetResult() interface{} {
	return p.Success
}

func queryUserFavoriteCountListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.FavoriteModelQueryUserFavoriteCountListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FavoriteModelService).QueryUserFavoriteCountList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryUserFavoriteCountListArgs:
		success, err := handler.(api.FavoriteModelService).QueryUserFavoriteCountList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryUserFavoriteCountListResult)
		realResult.Success = success
	}
	return nil
}
func newQueryUserFavoriteCountListArgs() interface{} {
	return &QueryUserFavoriteCountListArgs{}
}

func newQueryUserFavoriteCountListResult() interface{} {
	return &QueryUserFavoriteCountListResult{}
}

type QueryUserFavoriteCountListArgs struct {
	Req *api.FavoriteModelQueryUserFavoriteCountListRequest
}

func (p *QueryUserFavoriteCountListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.FavoriteModelQueryUserFavoriteCountListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryUserFavoriteCountListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryUserFavoriteCountListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryUserFavoriteCountListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryUserFavoriteCountListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryUserFavoriteCountListArgs) Unmarshal(in []byte) error {
	msg := new(api.FavoriteModelQueryUserFavoriteCountListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryUserFavoriteCountListArgs_Req_DEFAULT *api.FavoriteModelQueryUserFavoriteCountListRequest

func (p *QueryUserFavoriteCountListArgs) GetReq() *api.FavoriteModelQueryUserFavoriteCountListRequest {
	if !p.IsSetReq() {
		return QueryUserFavoriteCountListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryUserFavoriteCountListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryUserFavoriteCountListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryUserFavoriteCountListResult struct {
	Success *api.FavoriteModelQueryUserFavoriteCountListResponse
}

var QueryUserFavoriteCountListResult_Success_DEFAULT *api.FavoriteModelQueryUserFavoriteCountListResponse

func (p *QueryUserFavoriteCountListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.FavoriteModelQueryUserFavoriteCountListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryUserFavoriteCountListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryUserFavoriteCountListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryUserFavoriteCountListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryUserFavoriteCountListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryUserFavoriteCountListResult) Unmarshal(in []byte) error {
	msg := new(api.FavoriteModelQueryUserFavoriteCountListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryUserFavoriteCountListResult) GetSuccess() *api.FavoriteModelQueryUserFavoriteCountListResponse {
	if !p.IsSetSuccess() {
		return QueryUserFavoriteCountListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryUserFavoriteCountListResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.FavoriteModelQueryUserFavoriteCountListResponse)
}

func (p *QueryUserFavoriteCountListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryUserFavoriteCountListResult) GetResult() interface{} {
	return p.Success
}

func queryVideoFavoriteCountListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.FavoriteModelQueryVideoFavoriteCountListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FavoriteModelService).QueryVideoFavoriteCountList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryVideoFavoriteCountListArgs:
		success, err := handler.(api.FavoriteModelService).QueryVideoFavoriteCountList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryVideoFavoriteCountListResult)
		realResult.Success = success
	}
	return nil
}
func newQueryVideoFavoriteCountListArgs() interface{} {
	return &QueryVideoFavoriteCountListArgs{}
}

func newQueryVideoFavoriteCountListResult() interface{} {
	return &QueryVideoFavoriteCountListResult{}
}

type QueryVideoFavoriteCountListArgs struct {
	Req *api.FavoriteModelQueryVideoFavoriteCountListRequest
}

func (p *QueryVideoFavoriteCountListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.FavoriteModelQueryVideoFavoriteCountListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryVideoFavoriteCountListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryVideoFavoriteCountListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryVideoFavoriteCountListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryVideoFavoriteCountListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryVideoFavoriteCountListArgs) Unmarshal(in []byte) error {
	msg := new(api.FavoriteModelQueryVideoFavoriteCountListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryVideoFavoriteCountListArgs_Req_DEFAULT *api.FavoriteModelQueryVideoFavoriteCountListRequest

func (p *QueryVideoFavoriteCountListArgs) GetReq() *api.FavoriteModelQueryVideoFavoriteCountListRequest {
	if !p.IsSetReq() {
		return QueryVideoFavoriteCountListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryVideoFavoriteCountListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryVideoFavoriteCountListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryVideoFavoriteCountListResult struct {
	Success *api.FavoriteModelQueryVideoFavoriteCountListResponse
}

var QueryVideoFavoriteCountListResult_Success_DEFAULT *api.FavoriteModelQueryVideoFavoriteCountListResponse

func (p *QueryVideoFavoriteCountListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.FavoriteModelQueryVideoFavoriteCountListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryVideoFavoriteCountListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryVideoFavoriteCountListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryVideoFavoriteCountListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryVideoFavoriteCountListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryVideoFavoriteCountListResult) Unmarshal(in []byte) error {
	msg := new(api.FavoriteModelQueryVideoFavoriteCountListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryVideoFavoriteCountListResult) GetSuccess() *api.FavoriteModelQueryVideoFavoriteCountListResponse {
	if !p.IsSetSuccess() {
		return QueryVideoFavoriteCountListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryVideoFavoriteCountListResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.FavoriteModelQueryVideoFavoriteCountListResponse)
}

func (p *QueryVideoFavoriteCountListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryVideoFavoriteCountListResult) GetResult() interface{} {
	return p.Success
}

func queryVideoFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.FavoriteModelQueryVideoFavoriteCountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FavoriteModelService).QueryVideoFavoriteCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryVideoFavoriteCountArgs:
		success, err := handler.(api.FavoriteModelService).QueryVideoFavoriteCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryVideoFavoriteCountResult)
		realResult.Success = success
	}
	return nil
}
func newQueryVideoFavoriteCountArgs() interface{} {
	return &QueryVideoFavoriteCountArgs{}
}

func newQueryVideoFavoriteCountResult() interface{} {
	return &QueryVideoFavoriteCountResult{}
}

type QueryVideoFavoriteCountArgs struct {
	Req *api.FavoriteModelQueryVideoFavoriteCountRequest
}

func (p *QueryVideoFavoriteCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.FavoriteModelQueryVideoFavoriteCountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryVideoFavoriteCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryVideoFavoriteCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryVideoFavoriteCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryVideoFavoriteCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryVideoFavoriteCountArgs) Unmarshal(in []byte) error {
	msg := new(api.FavoriteModelQueryVideoFavoriteCountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryVideoFavoriteCountArgs_Req_DEFAULT *api.FavoriteModelQueryVideoFavoriteCountRequest

func (p *QueryVideoFavoriteCountArgs) GetReq() *api.FavoriteModelQueryVideoFavoriteCountRequest {
	if !p.IsSetReq() {
		return QueryVideoFavoriteCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryVideoFavoriteCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryVideoFavoriteCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryVideoFavoriteCountResult struct {
	Success *api.FavoriteModelQueryVideoFavoriteCountResponse
}

var QueryVideoFavoriteCountResult_Success_DEFAULT *api.FavoriteModelQueryVideoFavoriteCountResponse

func (p *QueryVideoFavoriteCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.FavoriteModelQueryVideoFavoriteCountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryVideoFavoriteCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryVideoFavoriteCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryVideoFavoriteCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryVideoFavoriteCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryVideoFavoriteCountResult) Unmarshal(in []byte) error {
	msg := new(api.FavoriteModelQueryVideoFavoriteCountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryVideoFavoriteCountResult) GetSuccess() *api.FavoriteModelQueryVideoFavoriteCountResponse {
	if !p.IsSetSuccess() {
		return QueryVideoFavoriteCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryVideoFavoriteCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.FavoriteModelQueryVideoFavoriteCountResponse)
}

func (p *QueryVideoFavoriteCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryVideoFavoriteCountResult) GetResult() interface{} {
	return p.Success
}

func queryIsUserFavoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.FavoriteModelQueryIsUserFavoriteListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FavoriteModelService).QueryIsUserFavoriteList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryIsUserFavoriteListArgs:
		success, err := handler.(api.FavoriteModelService).QueryIsUserFavoriteList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryIsUserFavoriteListResult)
		realResult.Success = success
	}
	return nil
}
func newQueryIsUserFavoriteListArgs() interface{} {
	return &QueryIsUserFavoriteListArgs{}
}

func newQueryIsUserFavoriteListResult() interface{} {
	return &QueryIsUserFavoriteListResult{}
}

type QueryIsUserFavoriteListArgs struct {
	Req *api.FavoriteModelQueryIsUserFavoriteListRequest
}

func (p *QueryIsUserFavoriteListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.FavoriteModelQueryIsUserFavoriteListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryIsUserFavoriteListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryIsUserFavoriteListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryIsUserFavoriteListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryIsUserFavoriteListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryIsUserFavoriteListArgs) Unmarshal(in []byte) error {
	msg := new(api.FavoriteModelQueryIsUserFavoriteListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryIsUserFavoriteListArgs_Req_DEFAULT *api.FavoriteModelQueryIsUserFavoriteListRequest

func (p *QueryIsUserFavoriteListArgs) GetReq() *api.FavoriteModelQueryIsUserFavoriteListRequest {
	if !p.IsSetReq() {
		return QueryIsUserFavoriteListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryIsUserFavoriteListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryIsUserFavoriteListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryIsUserFavoriteListResult struct {
	Success *api.FavoriteModelQueryIsUserFavoriteListResponse
}

var QueryIsUserFavoriteListResult_Success_DEFAULT *api.FavoriteModelQueryIsUserFavoriteListResponse

func (p *QueryIsUserFavoriteListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.FavoriteModelQueryIsUserFavoriteListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryIsUserFavoriteListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryIsUserFavoriteListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryIsUserFavoriteListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryIsUserFavoriteListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryIsUserFavoriteListResult) Unmarshal(in []byte) error {
	msg := new(api.FavoriteModelQueryIsUserFavoriteListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryIsUserFavoriteListResult) GetSuccess() *api.FavoriteModelQueryIsUserFavoriteListResponse {
	if !p.IsSetSuccess() {
		return QueryIsUserFavoriteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryIsUserFavoriteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.FavoriteModelQueryIsUserFavoriteListResponse)
}

func (p *QueryIsUserFavoriteListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryIsUserFavoriteListResult) GetResult() interface{} {
	return p.Success
}

func queryIsUserFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.FavoriteModelQueryIsUserFavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.FavoriteModelService).QueryIsUserFavorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryIsUserFavoriteArgs:
		success, err := handler.(api.FavoriteModelService).QueryIsUserFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryIsUserFavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newQueryIsUserFavoriteArgs() interface{} {
	return &QueryIsUserFavoriteArgs{}
}

func newQueryIsUserFavoriteResult() interface{} {
	return &QueryIsUserFavoriteResult{}
}

type QueryIsUserFavoriteArgs struct {
	Req *api.FavoriteModelQueryIsUserFavoriteRequest
}

func (p *QueryIsUserFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.FavoriteModelQueryIsUserFavoriteRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QueryIsUserFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QueryIsUserFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QueryIsUserFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryIsUserFavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryIsUserFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(api.FavoriteModelQueryIsUserFavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryIsUserFavoriteArgs_Req_DEFAULT *api.FavoriteModelQueryIsUserFavoriteRequest

func (p *QueryIsUserFavoriteArgs) GetReq() *api.FavoriteModelQueryIsUserFavoriteRequest {
	if !p.IsSetReq() {
		return QueryIsUserFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryIsUserFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryIsUserFavoriteArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryIsUserFavoriteResult struct {
	Success *api.FavoriteModelQueryIsUserFavoriteResponse
}

var QueryIsUserFavoriteResult_Success_DEFAULT *api.FavoriteModelQueryIsUserFavoriteResponse

func (p *QueryIsUserFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.FavoriteModelQueryIsUserFavoriteResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QueryIsUserFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QueryIsUserFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QueryIsUserFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryIsUserFavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryIsUserFavoriteResult) Unmarshal(in []byte) error {
	msg := new(api.FavoriteModelQueryIsUserFavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryIsUserFavoriteResult) GetSuccess() *api.FavoriteModelQueryIsUserFavoriteResponse {
	if !p.IsSetSuccess() {
		return QueryIsUserFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryIsUserFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.FavoriteModelQueryIsUserFavoriteResponse)
}

func (p *QueryIsUserFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryIsUserFavoriteResult) GetResult() interface{} {
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

func (p *kClient) QueryFavoriteList(ctx context.Context, Req *api.FavoriteModelQueryFavoriteListRequest) (r *api.FavoriteModelQueryFavoriteListResponse, err error) {
	var _args QueryFavoriteListArgs
	_args.Req = Req
	var _result QueryFavoriteListResult
	if err = p.c.Call(ctx, "queryFavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteAction(ctx context.Context, Req *api.FavoriteModelFavoriteActionRequest) (r *api.FavoriteModelFavoriteActionResponse, err error) {
	var _args FavoriteActionArgs
	_args.Req = Req
	var _result FavoriteActionResult
	if err = p.c.Call(ctx, "favoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryUserFavoriteCount(ctx context.Context, Req *api.FavoriteModelQueryUserFavoriteCountRequest) (r *api.FavoriteModelQueryUserFavoriteCountResponse, err error) {
	var _args QueryUserFavoriteCountArgs
	_args.Req = Req
	var _result QueryUserFavoriteCountResult
	if err = p.c.Call(ctx, "queryUserFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryUserFavoriteCountList(ctx context.Context, Req *api.FavoriteModelQueryUserFavoriteCountListRequest) (r *api.FavoriteModelQueryUserFavoriteCountListResponse, err error) {
	var _args QueryUserFavoriteCountListArgs
	_args.Req = Req
	var _result QueryUserFavoriteCountListResult
	if err = p.c.Call(ctx, "queryUserFavoriteCountList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryVideoFavoriteCountList(ctx context.Context, Req *api.FavoriteModelQueryVideoFavoriteCountListRequest) (r *api.FavoriteModelQueryVideoFavoriteCountListResponse, err error) {
	var _args QueryVideoFavoriteCountListArgs
	_args.Req = Req
	var _result QueryVideoFavoriteCountListResult
	if err = p.c.Call(ctx, "queryVideoFavoriteCountList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryVideoFavoriteCount(ctx context.Context, Req *api.FavoriteModelQueryVideoFavoriteCountRequest) (r *api.FavoriteModelQueryVideoFavoriteCountResponse, err error) {
	var _args QueryVideoFavoriteCountArgs
	_args.Req = Req
	var _result QueryVideoFavoriteCountResult
	if err = p.c.Call(ctx, "queryVideoFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryIsUserFavoriteList(ctx context.Context, Req *api.FavoriteModelQueryIsUserFavoriteListRequest) (r *api.FavoriteModelQueryIsUserFavoriteListResponse, err error) {
	var _args QueryIsUserFavoriteListArgs
	_args.Req = Req
	var _result QueryIsUserFavoriteListResult
	if err = p.c.Call(ctx, "queryIsUserFavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryIsUserFavorite(ctx context.Context, Req *api.FavoriteModelQueryIsUserFavoriteRequest) (r *api.FavoriteModelQueryIsUserFavoriteResponse, err error) {
	var _args QueryIsUserFavoriteArgs
	_args.Req = Req
	var _result QueryIsUserFavoriteResult
	if err = p.c.Call(ctx, "queryIsUserFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
