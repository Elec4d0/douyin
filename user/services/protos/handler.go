package protos

import (
	"context"
	"log"
	"time"
	"user/model"
	api "user/services/protos/kitex_gen/api"
	"user/utils"
	"user/utils/jwt"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *api.DouyinUserLoginRequest) (resp *api.DouyinUserLoginResponse, err error) {
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
	user, err := model.FindUserByName(name)
	if err != nil {
		resp.StatusCode = 1

		statusMsg := "用户名或密码错误！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1
		resp.Token = ""
		return
	}

	// check password
	if user.Password != utils.SHA256(password) {
		resp.StatusCode = 1

		statusMsg := "用户名或密码错误！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1
		resp.Token = ""
		return
	}

	//生成Token
	resp.Token, err = jwt.GenerateToken(user.ID, user.Name)
	if err != nil {
		resp.StatusCode = 1
		statusMsg := "Token生成失败！"
		resp.StatusMsg = &statusMsg
		resp.UserId = -1
		resp.Token = ""
		return
	}

	resp.StatusCode = 0
	statusMsg := "登录成功！"
	resp.StatusMsg = &statusMsg
	resp.UserId = user.ID

	return
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *api.DouyinUserRegisterRequest) (resp *api.DouyinUserRegisterResponse, err error) {
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
	err = model.FindName(name)
	if err == nil {
		resp.StatusCode = 1

		statusMsg := "用户名重复！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1
		resp.Token = ""
		return
	}

	// create user model
	user := &model.User{Name: name,
		Password:       utils.SHA256(password),
		FollowCount:    0,
		FollowerCount:  0,
		WorkCount:      0,
		TotalFavorited: 0,
		FavoriteCount:  0,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Token:          "",
	}

	id, err := model.CreateUser(user)
	if err != nil {
		resp.StatusCode = 1

		statusMsg := "注册失败！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1
		resp.Token = ""
		log.Fatal(err)
		return
	}

	//登录验证成功，生成Token
	resp.Token, err = jwt.GenerateToken(user.ID, user.Name)
	if err != nil {
		resp.StatusCode = 1
		statusMsg := "Token生成失败！"
		resp.StatusMsg = &statusMsg
		resp.UserId = -1
		resp.Token = ""
		return
	}

	resp.StatusCode = 0

	statusMsg := "注册成功！"
	resp.StatusMsg = &statusMsg

	resp.UserId = id

	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *api.DouyinUserRequest) (resp *api.DouyinUserResponse, err error) {
	// id that need to be searched
	searchID := req.UserId
	// token is belonged to user who need to search
	resp = new(api.DouyinUserResponse)

	// through token find id
	// need more----------------------
	// ownerID := 100010

	// find user info
	user, err := model.FindUserByID(searchID)
	if err != nil {
		resp.StatusCode = 1

		statusMsg := "用户信息错误！"
		resp.StatusMsg = &statusMsg

		resp.User = nil
		return
	}

	// get isFollow -----------------------------
	var isFollow bool
	isFollow = false

	resp.StatusCode = 0

	statusMsg := "查询成功！"
	resp.StatusMsg = &statusMsg

	resp.User = convertUserTable(user, isFollow)

	return
}
