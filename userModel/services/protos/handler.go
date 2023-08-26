package protos

import (
	"context"
	"log"
	"time"
	"userModel/model"
	api "userModel/services/protos/kitex_gen/api"
	"userModel/utils"
	"userModel/utils/idGeneration"
)

// UserModelServiceImpl implements the last service interface defined in the IDL.
type UserModelServiceImpl struct{}

// CreateBaseUser implements the UserModelServiceImpl interface.
func (s *UserModelServiceImpl) CreateBaseUser(ctx context.Context, req *api.DouyinUserCreateBaseUserRequest) (resp *api.DouyinUserCreateBaseUserResponse, err error) {
	// TODO: Your code here...
	resp = new(api.DouyinUserCreateBaseUserResponse)
	log.Println("handler: ", resp)
	name := req.Username
	password := req.Password

	id, err := idGeneration.GeneratelID()
	if err != nil {
		resp.StatusCode = 1

		statusMsg := "model层：用户创建失败, id生成失败！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1

		log.Fatal(err)
		return
	}

	user := &model.User{
		ID:              id,
		Name:            name,
		Password:        utils.SHA256(password),
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	id, err = model.CreateUser(user)

	if err != nil {
		resp.StatusCode = 1

		statusMsg := "model层：用户创建失败！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1

		log.Fatal(err)
		return
	}
	resp.StatusCode = 0

	statusMsg := "model层：用户创建成功！"
	resp.StatusMsg = &statusMsg

	resp.UserId = id
	return
}

// FindBaseUserByName implements the UserModelServiceImpl interface.
func (s *UserModelServiceImpl) FindBaseUserByName(ctx context.Context, req *api.DouyinUserFindBaseUserByNameRequest) (resp *api.DouyinUserFindBaseUserByNameResponse, err error) {
	// TODO: Your code here...
	resp = new(api.DouyinUserFindBaseUserByNameResponse)
	name := req.Username

	user, err := model.FindUserByName(name)

	if err != nil {
		resp.StatusCode = 1

		statusMsg := "model层：通过name的用户查找失败！"
		resp.StatusMsg = &statusMsg

		resp.BaseUser = nil

		log.Println(err)
		return
	}
	resp.StatusCode = 0

	statusMsg := "model层：通过name的用户查找成功！"
	resp.StatusMsg = &statusMsg

	resp.BaseUser = utils.ConvertUserTableToBaseUser(&user)
	return
}

// FindBaseUserById implements the UserModelServiceImpl interface.
func (s *UserModelServiceImpl) FindBaseUserById(ctx context.Context, req *api.DouyinUserFindBaseUserByIdRequest) (resp *api.DouyinUserFindBaseUserByIdResponse, err error) {
	// TODO: Your code here...
	resp = new(api.DouyinUserFindBaseUserByIdResponse)
	id := req.UserId

	user, err := model.FindUserByID(id)

	if err != nil {
		resp.StatusCode = 1

		statusMsg := "model层：通过id的用户查找失败！"
		resp.StatusMsg = &statusMsg

		resp.BaseUser = nil

		log.Println(err)
		return
	}
	resp.StatusCode = 0

	statusMsg := "model层：通过id的用户查找成功！"
	resp.StatusMsg = &statusMsg

	resp.BaseUser = utils.ConvertUserTableToBaseUser(&user)
	return
}

// FindBaseUserList implements the UserModelServiceImpl interface.
func (s *UserModelServiceImpl) FindBaseUserList(ctx context.Context, req *api.DouyinUserFindBaseUserListRequest) (resp *api.DouyinUserFindBaseUserListResponse, err error) {
	// TODO: Your code here...
	resp = new(api.DouyinUserFindBaseUserListResponse)

	authorId := req.AuthorId

	var users []*api.BaseUser

	userList, err := model.FindUserByIDs(authorId)

	for i := 0; i < len(authorId); i++ {
		users = append(users, utils.ConvertUserTableToBaseUser(userList[i]))
	}

	if err != nil {
		resp.StatusCode = 1

		statusMsg := "model层：通过id的用户List查找失败！"
		resp.StatusMsg = &statusMsg

		resp.BaseUser = users

		log.Println(err)
		return
	}
	resp.StatusCode = 0

	statusMsg := "model层：通过id的用户list查找成功！"
	resp.StatusMsg = &statusMsg

	resp.BaseUser = users

	return
}

// FindBaseUserPassword implements the UserModelServiceImpl interface.
func (s *UserModelServiceImpl) FindBaseUserPassword(ctx context.Context, req *api.DouyinUserFindBaseUserPasswordRequest) (resp *api.DouyinUserFindBaseUserPasswordResponse, err error) {
	// TODO: Your code here...
	resp = new(api.DouyinUserFindBaseUserPasswordResponse)
	userId := req.UserId
	var password string
	password, err = model.FindPasswordByID(userId)
	if err != nil {
		resp.StatusCode = 1

		statusMsg := "model层：通过id的用户密码查找失败！"
		resp.StatusMsg = &statusMsg

		resp.Password = ""
		log.Println(err)
		return
	}
	resp.StatusCode = 0

	statusMsg := "model层：通过id的用户密码查找成功！"
	resp.StatusMsg = &statusMsg

	resp.Password = password
	return
}

// FindIDByName implements the UserModelServiceImpl interface.
func (s *UserModelServiceImpl) FindIDByName(ctx context.Context, req *api.DouyinUserFindIdByNameRequest) (resp *api.DouyinUserFindIdByNameResponse, err error) {
	// TODO: Your code here...
	resp = new(api.DouyinUserFindIdByNameResponse)
	name := req.Name
	var id int64
	id, err = model.FindIDByName(name)
	if err != nil {
		resp.StatusCode = 1

		statusMsg := "model层：通过name的用户id查找失败！"
		resp.StatusMsg = &statusMsg

		resp.UserId = -1
		log.Println(err)
		return
	}
	resp.StatusCode = 0

	statusMsg := "model层：通过name的用户id查找成功！"
	resp.StatusMsg = &statusMsg

	resp.UserId = id
	return
}
