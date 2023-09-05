package protos

import (
	"context"
	"log"
	api "user/services/protos/kitex_gen/api"
	userModelServices "user/userModelAPI"
	"user/utils"
	"user/utils/jwt"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *api.DouyinUserLoginRequest) (resp *api.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	resp = new(api.DouyinUserLoginResponse)

	name := req.Username
	password := req.Password
	// empty
	if name == "" || password == "" {
		resp.StatusCode = 1

		statusMsg := "请输入非空的用户名和密码！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1
		resp.Token = ""
		return
	}

	// find id
	id, err := userModelServices.FindIDByName(name)
	if err != nil {
		resp.StatusCode = 1

		statusMsg := "用户名或密码错误！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1
		resp.Token = ""
		err = nil
		return
	}

	// check password
	psd, err := userModelServices.FindBaseUserPassword(id)
	if psd != utils.SHA256(password) {
		resp.StatusCode = 1

		statusMsg := "用户名或密码错误！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1
		resp.Token = ""
		return
	}

	// get token
	token, err := jwt.GenerateToken(id, name)
	if err != nil {
		resp.StatusCode = 1

		statusMsg := "token签发错误！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1
		resp.Token = ""
		return
	}

	resp.StatusCode = 0

	statusMsg := "登录成功！"
	resp.StatusMsg = &statusMsg

	resp.UserId = id
	resp.Token = token
	return
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *api.DouyinUserRegisterRequest) (resp *api.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(api.DouyinUserRegisterResponse)

	name := req.Username
	password := req.Password
	// empty
	if name == "" || password == "" {
		resp.StatusCode = 1

		statusMsg := "请输入非空的用户名和密码！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1
		resp.Token = ""
		return
	}

	// if there has the same name in database
	_, err = userModelServices.FindIDByName(name)
	if err == nil {
		resp.StatusCode = 1

		statusMsg := "用户名重复！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1
		resp.Token = ""
		return
	}

	id, err := userModelServices.CreateBaseUser(name, password)
	if err != nil {
		resp.StatusCode = 1

		statusMsg := "注册失败！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1
		resp.Token = ""
		log.Println(err)
		return
	}

	token, err := jwt.GenerateToken(id, name)
	if err != nil {
		resp.StatusCode = 1

		statusMsg := "token签发错误！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1
		resp.Token = ""
		return
	}
	resp.StatusCode = 0

	statusMsg := "注册成功！"
	resp.StatusMsg = &statusMsg

	resp.UserId = id
	resp.Token = token
	return
}
