// Code generated by Kitex v0.6.2. DO NOT EDIT.

package usermodelservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	api "userModel/services/protos/kitex_gen/api"
)

func serviceInfo() *kitex.ServiceInfo {
	return userModelServiceServiceInfo
}

var userModelServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserModelService"
	handlerType := (*api.UserModelService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateBaseUser":       kitex.NewMethodInfo(createBaseUserHandler, newCreateBaseUserArgs, newCreateBaseUserResult, false),
		"FindBaseUserByName":   kitex.NewMethodInfo(findBaseUserByNameHandler, newFindBaseUserByNameArgs, newFindBaseUserByNameResult, false),
		"FindBaseUserById":     kitex.NewMethodInfo(findBaseUserByIdHandler, newFindBaseUserByIdArgs, newFindBaseUserByIdResult, false),
		"FindBaseUserList":     kitex.NewMethodInfo(findBaseUserListHandler, newFindBaseUserListArgs, newFindBaseUserListResult, false),
		"FindBaseUserPassword": kitex.NewMethodInfo(findBaseUserPasswordHandler, newFindBaseUserPasswordArgs, newFindBaseUserPasswordResult, false),
		"FindIDByName":         kitex.NewMethodInfo(findIDByNameHandler, newFindIDByNameArgs, newFindIDByNameResult, false),
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

func createBaseUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinUserCreateBaseUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.UserModelService).CreateBaseUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateBaseUserArgs:
		success, err := handler.(api.UserModelService).CreateBaseUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateBaseUserResult)
		realResult.Success = success
	}
	return nil
}
func newCreateBaseUserArgs() interface{} {
	return &CreateBaseUserArgs{}
}

func newCreateBaseUserResult() interface{} {
	return &CreateBaseUserResult{}
}

type CreateBaseUserArgs struct {
	Req *api.DouyinUserCreateBaseUserRequest
}

func (p *CreateBaseUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinUserCreateBaseUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateBaseUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateBaseUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateBaseUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateBaseUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateBaseUserArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserCreateBaseUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateBaseUserArgs_Req_DEFAULT *api.DouyinUserCreateBaseUserRequest

func (p *CreateBaseUserArgs) GetReq() *api.DouyinUserCreateBaseUserRequest {
	if !p.IsSetReq() {
		return CreateBaseUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateBaseUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateBaseUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateBaseUserResult struct {
	Success *api.DouyinUserCreateBaseUserResponse
}

var CreateBaseUserResult_Success_DEFAULT *api.DouyinUserCreateBaseUserResponse

func (p *CreateBaseUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinUserCreateBaseUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateBaseUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateBaseUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateBaseUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateBaseUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateBaseUserResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserCreateBaseUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateBaseUserResult) GetSuccess() *api.DouyinUserCreateBaseUserResponse {
	if !p.IsSetSuccess() {
		return CreateBaseUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateBaseUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinUserCreateBaseUserResponse)
}

func (p *CreateBaseUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateBaseUserResult) GetResult() interface{} {
	return p.Success
}

func findBaseUserByNameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinUserFindBaseUserByNameRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.UserModelService).FindBaseUserByName(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FindBaseUserByNameArgs:
		success, err := handler.(api.UserModelService).FindBaseUserByName(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FindBaseUserByNameResult)
		realResult.Success = success
	}
	return nil
}
func newFindBaseUserByNameArgs() interface{} {
	return &FindBaseUserByNameArgs{}
}

func newFindBaseUserByNameResult() interface{} {
	return &FindBaseUserByNameResult{}
}

type FindBaseUserByNameArgs struct {
	Req *api.DouyinUserFindBaseUserByNameRequest
}

func (p *FindBaseUserByNameArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinUserFindBaseUserByNameRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FindBaseUserByNameArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FindBaseUserByNameArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FindBaseUserByNameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FindBaseUserByNameArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FindBaseUserByNameArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserFindBaseUserByNameRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FindBaseUserByNameArgs_Req_DEFAULT *api.DouyinUserFindBaseUserByNameRequest

func (p *FindBaseUserByNameArgs) GetReq() *api.DouyinUserFindBaseUserByNameRequest {
	if !p.IsSetReq() {
		return FindBaseUserByNameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FindBaseUserByNameArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FindBaseUserByNameArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FindBaseUserByNameResult struct {
	Success *api.DouyinUserFindBaseUserByNameResponse
}

var FindBaseUserByNameResult_Success_DEFAULT *api.DouyinUserFindBaseUserByNameResponse

func (p *FindBaseUserByNameResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinUserFindBaseUserByNameResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FindBaseUserByNameResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FindBaseUserByNameResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FindBaseUserByNameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FindBaseUserByNameResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FindBaseUserByNameResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserFindBaseUserByNameResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FindBaseUserByNameResult) GetSuccess() *api.DouyinUserFindBaseUserByNameResponse {
	if !p.IsSetSuccess() {
		return FindBaseUserByNameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FindBaseUserByNameResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinUserFindBaseUserByNameResponse)
}

func (p *FindBaseUserByNameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FindBaseUserByNameResult) GetResult() interface{} {
	return p.Success
}

func findBaseUserByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinUserFindBaseUserByIdRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.UserModelService).FindBaseUserById(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FindBaseUserByIdArgs:
		success, err := handler.(api.UserModelService).FindBaseUserById(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FindBaseUserByIdResult)
		realResult.Success = success
	}
	return nil
}
func newFindBaseUserByIdArgs() interface{} {
	return &FindBaseUserByIdArgs{}
}

func newFindBaseUserByIdResult() interface{} {
	return &FindBaseUserByIdResult{}
}

type FindBaseUserByIdArgs struct {
	Req *api.DouyinUserFindBaseUserByIdRequest
}

func (p *FindBaseUserByIdArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinUserFindBaseUserByIdRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FindBaseUserByIdArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FindBaseUserByIdArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FindBaseUserByIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FindBaseUserByIdArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FindBaseUserByIdArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserFindBaseUserByIdRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FindBaseUserByIdArgs_Req_DEFAULT *api.DouyinUserFindBaseUserByIdRequest

func (p *FindBaseUserByIdArgs) GetReq() *api.DouyinUserFindBaseUserByIdRequest {
	if !p.IsSetReq() {
		return FindBaseUserByIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FindBaseUserByIdArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FindBaseUserByIdArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FindBaseUserByIdResult struct {
	Success *api.DouyinUserFindBaseUserByIdResponse
}

var FindBaseUserByIdResult_Success_DEFAULT *api.DouyinUserFindBaseUserByIdResponse

func (p *FindBaseUserByIdResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinUserFindBaseUserByIdResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FindBaseUserByIdResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FindBaseUserByIdResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FindBaseUserByIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FindBaseUserByIdResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FindBaseUserByIdResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserFindBaseUserByIdResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FindBaseUserByIdResult) GetSuccess() *api.DouyinUserFindBaseUserByIdResponse {
	if !p.IsSetSuccess() {
		return FindBaseUserByIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FindBaseUserByIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinUserFindBaseUserByIdResponse)
}

func (p *FindBaseUserByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FindBaseUserByIdResult) GetResult() interface{} {
	return p.Success
}

func findBaseUserListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinUserFindBaseUserListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.UserModelService).FindBaseUserList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FindBaseUserListArgs:
		success, err := handler.(api.UserModelService).FindBaseUserList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FindBaseUserListResult)
		realResult.Success = success
	}
	return nil
}
func newFindBaseUserListArgs() interface{} {
	return &FindBaseUserListArgs{}
}

func newFindBaseUserListResult() interface{} {
	return &FindBaseUserListResult{}
}

type FindBaseUserListArgs struct {
	Req *api.DouyinUserFindBaseUserListRequest
}

func (p *FindBaseUserListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinUserFindBaseUserListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FindBaseUserListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FindBaseUserListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FindBaseUserListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FindBaseUserListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FindBaseUserListArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserFindBaseUserListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FindBaseUserListArgs_Req_DEFAULT *api.DouyinUserFindBaseUserListRequest

func (p *FindBaseUserListArgs) GetReq() *api.DouyinUserFindBaseUserListRequest {
	if !p.IsSetReq() {
		return FindBaseUserListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FindBaseUserListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FindBaseUserListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FindBaseUserListResult struct {
	Success *api.DouyinUserFindBaseUserListResponse
}

var FindBaseUserListResult_Success_DEFAULT *api.DouyinUserFindBaseUserListResponse

func (p *FindBaseUserListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinUserFindBaseUserListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FindBaseUserListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FindBaseUserListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FindBaseUserListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FindBaseUserListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FindBaseUserListResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserFindBaseUserListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FindBaseUserListResult) GetSuccess() *api.DouyinUserFindBaseUserListResponse {
	if !p.IsSetSuccess() {
		return FindBaseUserListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FindBaseUserListResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinUserFindBaseUserListResponse)
}

func (p *FindBaseUserListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FindBaseUserListResult) GetResult() interface{} {
	return p.Success
}

func findBaseUserPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinUserFindBaseUserPasswordRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.UserModelService).FindBaseUserPassword(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FindBaseUserPasswordArgs:
		success, err := handler.(api.UserModelService).FindBaseUserPassword(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FindBaseUserPasswordResult)
		realResult.Success = success
	}
	return nil
}
func newFindBaseUserPasswordArgs() interface{} {
	return &FindBaseUserPasswordArgs{}
}

func newFindBaseUserPasswordResult() interface{} {
	return &FindBaseUserPasswordResult{}
}

type FindBaseUserPasswordArgs struct {
	Req *api.DouyinUserFindBaseUserPasswordRequest
}

func (p *FindBaseUserPasswordArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinUserFindBaseUserPasswordRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FindBaseUserPasswordArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FindBaseUserPasswordArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FindBaseUserPasswordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FindBaseUserPasswordArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FindBaseUserPasswordArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserFindBaseUserPasswordRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FindBaseUserPasswordArgs_Req_DEFAULT *api.DouyinUserFindBaseUserPasswordRequest

func (p *FindBaseUserPasswordArgs) GetReq() *api.DouyinUserFindBaseUserPasswordRequest {
	if !p.IsSetReq() {
		return FindBaseUserPasswordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FindBaseUserPasswordArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FindBaseUserPasswordArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FindBaseUserPasswordResult struct {
	Success *api.DouyinUserFindBaseUserPasswordResponse
}

var FindBaseUserPasswordResult_Success_DEFAULT *api.DouyinUserFindBaseUserPasswordResponse

func (p *FindBaseUserPasswordResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinUserFindBaseUserPasswordResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FindBaseUserPasswordResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FindBaseUserPasswordResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FindBaseUserPasswordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FindBaseUserPasswordResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FindBaseUserPasswordResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserFindBaseUserPasswordResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FindBaseUserPasswordResult) GetSuccess() *api.DouyinUserFindBaseUserPasswordResponse {
	if !p.IsSetSuccess() {
		return FindBaseUserPasswordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FindBaseUserPasswordResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinUserFindBaseUserPasswordResponse)
}

func (p *FindBaseUserPasswordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FindBaseUserPasswordResult) GetResult() interface{} {
	return p.Success
}

func findIDByNameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DouyinUserFindIdByNameRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.UserModelService).FindIDByName(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FindIDByNameArgs:
		success, err := handler.(api.UserModelService).FindIDByName(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FindIDByNameResult)
		realResult.Success = success
	}
	return nil
}
func newFindIDByNameArgs() interface{} {
	return &FindIDByNameArgs{}
}

func newFindIDByNameResult() interface{} {
	return &FindIDByNameResult{}
}

type FindIDByNameArgs struct {
	Req *api.DouyinUserFindIdByNameRequest
}

func (p *FindIDByNameArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DouyinUserFindIdByNameRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FindIDByNameArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FindIDByNameArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FindIDByNameArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FindIDByNameArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FindIDByNameArgs) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserFindIdByNameRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FindIDByNameArgs_Req_DEFAULT *api.DouyinUserFindIdByNameRequest

func (p *FindIDByNameArgs) GetReq() *api.DouyinUserFindIdByNameRequest {
	if !p.IsSetReq() {
		return FindIDByNameArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FindIDByNameArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FindIDByNameArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FindIDByNameResult struct {
	Success *api.DouyinUserFindIdByNameResponse
}

var FindIDByNameResult_Success_DEFAULT *api.DouyinUserFindIdByNameResponse

func (p *FindIDByNameResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DouyinUserFindIdByNameResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FindIDByNameResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FindIDByNameResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FindIDByNameResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FindIDByNameResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FindIDByNameResult) Unmarshal(in []byte) error {
	msg := new(api.DouyinUserFindIdByNameResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FindIDByNameResult) GetSuccess() *api.DouyinUserFindIdByNameResponse {
	if !p.IsSetSuccess() {
		return FindIDByNameResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FindIDByNameResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DouyinUserFindIdByNameResponse)
}

func (p *FindIDByNameResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FindIDByNameResult) GetResult() interface{} {
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

func (p *kClient) CreateBaseUser(ctx context.Context, Req *api.DouyinUserCreateBaseUserRequest) (r *api.DouyinUserCreateBaseUserResponse, err error) {
	var _args CreateBaseUserArgs
	_args.Req = Req
	var _result CreateBaseUserResult
	if err = p.c.Call(ctx, "CreateBaseUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FindBaseUserByName(ctx context.Context, Req *api.DouyinUserFindBaseUserByNameRequest) (r *api.DouyinUserFindBaseUserByNameResponse, err error) {
	var _args FindBaseUserByNameArgs
	_args.Req = Req
	var _result FindBaseUserByNameResult
	if err = p.c.Call(ctx, "FindBaseUserByName", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FindBaseUserById(ctx context.Context, Req *api.DouyinUserFindBaseUserByIdRequest) (r *api.DouyinUserFindBaseUserByIdResponse, err error) {
	var _args FindBaseUserByIdArgs
	_args.Req = Req
	var _result FindBaseUserByIdResult
	if err = p.c.Call(ctx, "FindBaseUserById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FindBaseUserList(ctx context.Context, Req *api.DouyinUserFindBaseUserListRequest) (r *api.DouyinUserFindBaseUserListResponse, err error) {
	var _args FindBaseUserListArgs
	_args.Req = Req
	var _result FindBaseUserListResult
	if err = p.c.Call(ctx, "FindBaseUserList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FindBaseUserPassword(ctx context.Context, Req *api.DouyinUserFindBaseUserPasswordRequest) (r *api.DouyinUserFindBaseUserPasswordResponse, err error) {
	var _args FindBaseUserPasswordArgs
	_args.Req = Req
	var _result FindBaseUserPasswordResult
	if err = p.c.Call(ctx, "FindBaseUserPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FindIDByName(ctx context.Context, Req *api.DouyinUserFindIdByNameRequest) (r *api.DouyinUserFindIdByNameResponse, err error) {
	var _args FindIDByNameArgs
	_args.Req = Req
	var _result FindIDByNameResult
	if err = p.c.Call(ctx, "FindIDByName", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
