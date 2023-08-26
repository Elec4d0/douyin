// Code generated by Kitex v0.6.2. DO NOT EDIT.

package userservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	api "gateway/microService/user/api"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*api.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"UserLogin":    kitex.NewMethodInfo(userLoginHandler, newUserLoginArgs, newUserLoginResult, false),
		"UserRegister": kitex.NewMethodInfo(userRegisterHandler, newUserRegisterArgs, newUserRegisterResult, false),
		"UserInfo":     kitex.NewMethodInfo(userInfoHandler, newUserInfoArgs, newUserInfoResult, false),
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

func userLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinUserLoginRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.UserService).UserLogin(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserLoginArgs:
		success, err := handler.(api.UserService).UserLogin(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserLoginResult)
		realResult.Success = success
	}
	return nil
}
func newUserLoginArgs() interface{} {
	return &UserLoginArgs{}
}

func newUserLoginResult() interface{} {
	return &UserLoginResult{}
}

type UserLoginArgs struct {
	Req *api.DouyinUserLoginRequest
}

func (p *UserLoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinUserLoginRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserLoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserLoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserLoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserLoginArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserLoginArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserLoginRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserLoginArgs_Req_DEFAULT *api.DouyinUserLoginRequest

func (p *UserLoginArgs) GetReq() *api.DouyinUserLoginRequest {
	if !p.IsSetReq() {
		return UserLoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserLoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserLoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UserLoginResult struct {
	Success *api.DouyinUserLoginResponse
}

var UserLoginResult_Success_DEFAULT *api.DouyinUserLoginResponse

func (p *UserLoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinUserLoginResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserLoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserLoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserLoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserLoginResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserLoginResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserLoginResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserLoginResult) GetSuccess() *api.DouyinUserLoginResponse {
	if !p.IsSetSuccess() {
		return UserLoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserLoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinUserLoginResponse)
}

func (p *UserLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserLoginResult) GetResult() interface{} {
	return p.Success
}

func userRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinUserRegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.UserService).UserRegister(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserRegisterArgs:
		success, err := handler.(api.UserService).UserRegister(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserRegisterResult)
		realResult.Success = success
	}
	return nil
}
func newUserRegisterArgs() interface{} {
	return &UserRegisterArgs{}
}

func newUserRegisterResult() interface{} {
	return &UserRegisterResult{}
}

type UserRegisterArgs struct {
	Req *api.DouyinUserRegisterRequest
}

func (p *UserRegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinUserRegisterRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserRegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserRegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserRegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserRegisterArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserRegisterArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserRegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserRegisterArgs_Req_DEFAULT *api.DouyinUserRegisterRequest

func (p *UserRegisterArgs) GetReq() *api.DouyinUserRegisterRequest {
	if !p.IsSetReq() {
		return UserRegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserRegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserRegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UserRegisterResult struct {
	Success *api.DouyinUserRegisterResponse
}

var UserRegisterResult_Success_DEFAULT *api.DouyinUserRegisterResponse

func (p *UserRegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinUserRegisterResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserRegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserRegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserRegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserRegisterResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserRegisterResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserRegisterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserRegisterResult) GetSuccess() *api.DouyinUserRegisterResponse {
	if !p.IsSetSuccess() {
		return UserRegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserRegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinUserRegisterResponse)
}

func (p *UserRegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserRegisterResult) GetResult() interface{} {
	return p.Success
}

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.UserService).UserInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserInfoArgs:
		success, err := handler.(api.UserService).UserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserInfoResult)
		realResult.Success = success
	}
	return nil
}
func newUserInfoArgs() interface{} {
	return &UserInfoArgs{}
}

func newUserInfoResult() interface{} {
	return &UserInfoResult{}
}

type UserInfoArgs struct {
	Req *api.DouyinUserRequest
}

func (p *UserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserInfoArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserInfoArgs_Req_DEFAULT *api.DouyinUserRequest

func (p *UserInfoArgs) GetReq() *api.DouyinUserRequest {
	if !p.IsSetReq() {
		return UserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UserInfoResult struct {
	Success *api.DouyinUserResponse
}

var UserInfoResult_Success_DEFAULT *api.DouyinUserResponse

func (p *UserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserInfoResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserInfoResult) GetSuccess() *api.DouyinUserResponse {
	if !p.IsSetSuccess() {
		return UserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinUserResponse)
}

func (p *UserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserInfoResult) GetResult() interface{} {
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

func (p *kClient) UserLogin(ctx context.Context, Req *api.DouyinUserLoginRequest) (r *api.DouyinUserLoginResponse, err error) {
	var _args UserLoginArgs
	_args.Req = Req
	var _result UserLoginResult
	if err = p.c.Call(ctx, "UserLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserRegister(ctx context.Context, Req *api.DouyinUserRegisterRequest) (r *api.DouyinUserRegisterResponse, err error) {
	var _args UserRegisterArgs
	_args.Req = Req
	var _result UserRegisterResult
	if err = p.c.Call(ctx, "UserRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfo(ctx context.Context, Req *api.DouyinUserRequest) (r *api.DouyinUserResponse, err error) {
	var _args UserInfoArgs
	_args.Req = Req
	var _result UserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
