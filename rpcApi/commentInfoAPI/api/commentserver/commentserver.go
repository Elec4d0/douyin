// Code generated by Kitex v0.6.2. DO NOT EDIT.

package commentserver

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	api "rpcApi/commentInfoAPI/api"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentServerServiceInfo
}

var commentServerServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CommentServer"
	handlerType := (*api.CommentServer)(nil)
	methods := map[string]kitex.MethodInfo{
		"CommentAction":   kitex.NewMethodInfo(commentActionHandler, newCommentActionArgs, newCommentActionResult, false),
		"CommentList":     kitex.NewMethodInfo(commentListHandler, newCommentListArgs, newCommentListResult, false),
		"CommentCount":    kitex.NewMethodInfo(commentCountHandler, newCommentCountArgs, newCommentCountResult, false),
		"CommentAllCount": kitex.NewMethodInfo(commentAllCountHandler, newCommentAllCountArgs, newCommentAllCountResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "comment",
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

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinCommentActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.CommentServer).CommentAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CommentActionArgs:
		success, err := handler.(api.CommentServer).CommentAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CommentActionResult)
		realResult.Success = success
	}
	return nil
}
func newCommentActionArgs() interface{} {
	return &CommentActionArgs{}
}

func newCommentActionResult() interface{} {
	return &CommentActionResult{}
}

type CommentActionArgs struct {
	Req *api.DouyinCommentActionRequest
}

func (p *CommentActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinCommentActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CommentActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CommentActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CommentActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CommentActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CommentActionArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinCommentActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CommentActionArgs_Req_DEFAULT *api.DouyinCommentActionRequest

func (p *CommentActionArgs) GetReq() *api.DouyinCommentActionRequest {
	if !p.IsSetReq() {
		return CommentActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CommentActionArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CommentActionArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CommentActionResult struct {
	Success *api.DouyinCommentActionResponse
}

var CommentActionResult_Success_DEFAULT *api.DouyinCommentActionResponse

func (p *CommentActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinCommentActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CommentActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CommentActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CommentActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CommentActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CommentActionResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinCommentActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CommentActionResult) GetSuccess() *api.DouyinCommentActionResponse {
	if !p.IsSetSuccess() {
		return CommentActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CommentActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinCommentActionResponse)
}

func (p *CommentActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CommentActionResult) GetResult() interface{} {
	return p.Success
}

func commentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinCommentListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.CommentServer).CommentList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CommentListArgs:
		success, err := handler.(api.CommentServer).CommentList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CommentListResult)
		realResult.Success = success
	}
	return nil
}
func newCommentListArgs() interface{} {
	return &CommentListArgs{}
}

func newCommentListResult() interface{} {
	return &CommentListResult{}
}

type CommentListArgs struct {
	Req *api.DouyinCommentListRequest
}

func (p *CommentListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinCommentListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CommentListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CommentListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CommentListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CommentListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CommentListArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinCommentListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CommentListArgs_Req_DEFAULT *api.DouyinCommentListRequest

func (p *CommentListArgs) GetReq() *api.DouyinCommentListRequest {
	if !p.IsSetReq() {
		return CommentListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CommentListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CommentListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CommentListResult struct {
	Success *api.DouyinCommentListResponse
}

var CommentListResult_Success_DEFAULT *api.DouyinCommentListResponse

func (p *CommentListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinCommentListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CommentListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CommentListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CommentListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CommentListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CommentListResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinCommentListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CommentListResult) GetSuccess() *api.DouyinCommentListResponse {
	if !p.IsSetSuccess() {
		return CommentListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CommentListResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinCommentListResponse)
}

func (p *CommentListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CommentListResult) GetResult() interface{} {
	return p.Success
}

func commentCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinCommentserverCommentcountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.CommentServer).CommentCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CommentCountArgs:
		success, err := handler.(api.CommentServer).CommentCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CommentCountResult)
		realResult.Success = success
	}
	return nil
}
func newCommentCountArgs() interface{} {
	return &CommentCountArgs{}
}

func newCommentCountResult() interface{} {
	return &CommentCountResult{}
}

type CommentCountArgs struct {
	Req *api.DouyinCommentserverCommentcountRequest
}

func (p *CommentCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinCommentserverCommentcountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CommentCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CommentCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CommentCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CommentCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CommentCountArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinCommentserverCommentcountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CommentCountArgs_Req_DEFAULT *api.DouyinCommentserverCommentcountRequest

func (p *CommentCountArgs) GetReq() *api.DouyinCommentserverCommentcountRequest {
	if !p.IsSetReq() {
		return CommentCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CommentCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CommentCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CommentCountResult struct {
	Success *api.DouyinCommentserverCommentcountResponse
}

var CommentCountResult_Success_DEFAULT *api.DouyinCommentserverCommentcountResponse

func (p *CommentCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinCommentserverCommentcountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CommentCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CommentCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CommentCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CommentCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CommentCountResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinCommentserverCommentcountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CommentCountResult) GetSuccess() *api.DouyinCommentserverCommentcountResponse {
	if !p.IsSetSuccess() {
		return CommentCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CommentCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinCommentserverCommentcountResponse)
}

func (p *CommentCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CommentCountResult) GetResult() interface{} {
	return p.Success
}

func commentAllCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinCommentserverCommentallcountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.CommentServer).CommentAllCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CommentAllCountArgs:
		success, err := handler.(api.CommentServer).CommentAllCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CommentAllCountResult)
		realResult.Success = success
	}
	return nil
}
func newCommentAllCountArgs() interface{} {
	return &CommentAllCountArgs{}
}

func newCommentAllCountResult() interface{} {
	return &CommentAllCountResult{}
}

type CommentAllCountArgs struct {
	Req *api.DouyinCommentserverCommentallcountRequest
}

func (p *CommentAllCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinCommentserverCommentallcountRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CommentAllCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CommentAllCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CommentAllCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CommentAllCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CommentAllCountArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinCommentserverCommentallcountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CommentAllCountArgs_Req_DEFAULT *api.DouyinCommentserverCommentallcountRequest

func (p *CommentAllCountArgs) GetReq() *api.DouyinCommentserverCommentallcountRequest {
	if !p.IsSetReq() {
		return CommentAllCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CommentAllCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CommentAllCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CommentAllCountResult struct {
	Success *api.DouyinCommentserverCommentallcountResponse
}

var CommentAllCountResult_Success_DEFAULT *api.DouyinCommentserverCommentallcountResponse

func (p *CommentAllCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinCommentserverCommentallcountResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CommentAllCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CommentAllCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CommentAllCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CommentAllCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CommentAllCountResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinCommentserverCommentallcountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CommentAllCountResult) GetSuccess() *api.DouyinCommentserverCommentallcountResponse {
	if !p.IsSetSuccess() {
		return CommentAllCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CommentAllCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinCommentserverCommentallcountResponse)
}

func (p *CommentAllCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CommentAllCountResult) GetResult() interface{} {
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

func (p *kClient) CommentAction(ctx context.Context, Req *api.DouyinCommentActionRequest) (r *api.DouyinCommentActionResponse, err error) {
	var _args CommentActionArgs
	_args.Req = Req
	var _result CommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentList(ctx context.Context, Req *api.DouyinCommentListRequest) (r *api.DouyinCommentListResponse, err error) {
	var _args CommentListArgs
	_args.Req = Req
	var _result CommentListResult
	if err = p.c.Call(ctx, "CommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentCount(ctx context.Context, Req *api.DouyinCommentserverCommentcountRequest) (r *api.DouyinCommentserverCommentcountResponse, err error) {
	var _args CommentCountArgs
	_args.Req = Req
	var _result CommentCountResult
	if err = p.c.Call(ctx, "CommentCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentAllCount(ctx context.Context, Req *api.DouyinCommentserverCommentallcountRequest) (r *api.DouyinCommentserverCommentallcountResponse, err error) {
	var _args CommentAllCountArgs
	_args.Req = Req
	var _result CommentAllCountResult
	if err = p.c.Call(ctx, "CommentAllCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
