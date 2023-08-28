package query

import (
	"fmt"
	"userInfo/favoriteModel"
	"userInfo/services/protos/kitex_gen/api"
	userModelServices "userInfo/userModelAPI"
	"userInfo/utils/redis"
	videoModelServices "userInfo/videoModel"
)

func GetWorkCount(user_id int64) int64 {
	count, err := videoModelServices.QueryAuthorWorkCount(user_id)
	if err != nil {

		return 0
	}
	return int64(count)
}

func GetFavoriteInfo(user_id int64) (int64, int64) {
	likeCount, totalFavorited, err := favoriteModel.QueryUserFavoriteCount(user_id)
	if err != nil {
		return 0, 0
	}
	return likeCount, totalFavorited
}

func GetRelationInfo(user_id int64, search_id int64) (int64, int64, bool) {
	return 0, 0, false
}

func GetWorkCountList(user_id []int64) []int64 {
	count, _ := videoModelServices.QueryAuthorWorkCountList(user_id)
	//var workCount = make([]int64, len(user_id))
	//if count != nil {
	//	for i := 0; i < len(count); i++ {
	//		workCount = append(workCount, int64(count[i]))
	//	}
	//}
	var workCount []int64
	for i := 0; i < len(count); i++ {
		workCount = append(workCount, int64(count[i]))
	}
	return workCount
}

func GetFavoriteInfoList(user_id []int64) ([]int64, []int64) {
	likeCount, totalFavorited, _ := favoriteModel.BatchQueryUserFavoriteCount(user_id)

	return likeCount, totalFavorited
}

func GetRelationInfoList(user_id int64, search_id []int64) ([]int64, []int64, []bool) {
	return nil, nil, nil
}

func QueryUserInfo(userId int64, searchId int64) (*api.FullUser, error) {
	var fullUser *api.FullUser
	// ID是否存在以及是否过期
	isUserExist := redis.CheckUserExist(searchId)
	if isUserExist {
		//缓存存在，查询缓存和数据库的isFollow
		user, err := redis.FindUserRedis(searchId)
		// 查询成功
		if err == nil {
			fmt.Println("cache查询！！")
			// 查询成功， 判断是否nil
			if user == nil {
				fullUser = nil
				return nil, nil
			}

			_, _, isFollow := GetRelationInfo(userId, searchId)

			fullUser = &api.FullUser{
				Id:              user.Id,
				Name:            user.Name,
				Avatar:          user.Avatar,
				BackgroundImage: user.BackgroundImage,
				Signature:       user.Signature,
				WorkCount:       user.WorkCount,
				TotalFavorited:  user.TotalFavorited,
				FavoriteCount:   user.FavoriteCount,
				FollowerCount:   user.FollowerCount,
				FollowCount:     user.FollowCount,
				IsFollow:        isFollow,
			}

			return fullUser, nil
		}

	}

	baseUser, _ := userModelServices.FindBaseUserById(searchId)
	fmt.Println("model查询！！")
	if baseUser == nil {
		defer func() {
			_ = redis.InsertNilUserRedis(searchId)
		}()
		return nil, nil
	}

	followCount, followerCount, isFollow := GetRelationInfo(userId, searchId)

	favoriteCount, totalFavorited := GetFavoriteInfo(searchId)

	//获取视频点赞数
	workCount := GetWorkCount(searchId)

	id := baseUser.Id
	name := baseUser.Name
	avatar := baseUser.Avatar
	backgroundImage := baseUser.BackgroundImage
	signature := baseUser.Signature

	fullUser = &api.FullUser{
		Id:              id,
		Name:            name,
		Avatar:          avatar,
		BackgroundImage: backgroundImage,
		Signature:       signature,
		WorkCount:       &workCount,
		TotalFavorited:  &totalFavorited,
		FavoriteCount:   &favoriteCount,
		FollowerCount:   &followerCount,
		FollowCount:     &followCount,
		IsFollow:        isFollow,
	}

	defer func() {
		_ = redis.InsertUserRedis(fullUser)
	}()

	return fullUser, nil
}

func QueryUserListInfo(userId int64, searchId []int64) ([]*api.FullUser, error) {
	isUserListExist, flag := redis.CheckUserListExist(searchId)
	var fullUser []*api.FullUser
	var ok bool
	if flag == 2 {
		fullUser, ok = QueryUserListByCache(userId, searchId)
	} else if flag == 1 {
		fullUser, ok = QueryUserListByMix(userId, searchId, isUserListExist)
	}

	if ok {
		return fullUser, nil
	}

	fullUser, ok = QueryUserListByModel(userId, searchId)

	if ok {
		return fullUser, nil
	}

	return make([]*api.FullUser, len(searchId)), nil
}

func QueryUserListByModel(userId int64, searchId []int64) ([]*api.FullUser, bool) {
	fmt.Println("find list user_id model: ", userId)

	baseUser, _ := userModelServices.FindBaseUserList(searchId)
	// followCount, followerCount, isFollow := GetRelationInfoList(userId, searchId)
	followCount1 := int64(0)
	followerCount1 := int64(0)

	//获取用户与视频的喜好关系
	favoriteCount, totalFavorited := GetFavoriteInfoList(searchId)

	//获取视频点赞数
	workCount := GetWorkCountList(searchId)

	var fullUser []*api.FullUser

	for i := 0; i < len(searchId); i++ {
		if baseUser[i] == nil {
			fullUser = append(fullUser, nil)
		} else {
			if baseUser[i].Id == 0 {
				fullUser = append(fullUser, nil)
				continue
			}

			fullUser = append(fullUser,
				&api.FullUser{
					Id:              baseUser[i].Id,
					Name:            baseUser[i].Name,
					Avatar:          baseUser[i].Avatar,
					BackgroundImage: baseUser[i].Avatar,
					Signature:       baseUser[i].Avatar,
					WorkCount:       &workCount[i],
					TotalFavorited:  &totalFavorited[i],
					FavoriteCount:   &favoriteCount[i],
					FollowerCount:   &followerCount1,
					FollowCount:     &followCount1,
					IsFollow:        false,
				})

		}
	}
	defer func() {
		_ = redis.InsertUserListRedis(fullUser)
	}()

	return fullUser, true
}

func QueryUserListByCache(userId int64, searchId []int64) ([]*api.FullUser, bool) {
	fmt.Println("find list user_id cache: ", userId)
	var fullUser []*api.FullUser
	var user []*redis.User

	user, err := redis.FindUserListRedis(searchId)
	_, _, isFollow := GetRelationInfoList(userId, searchId)
	isFollow = make([]bool, len(searchId))

	if err != nil {
		return make([]*api.FullUser, len(searchId)), false
	}

	for i := 0; i < len(searchId); i++ {
		if user[i] == nil {
			fullUser = append(fullUser, nil)
		} else {
			fullUser = append(fullUser,
				&api.FullUser{
					Id:              user[i].Id,
					Name:            user[i].Name,
					Avatar:          user[i].Avatar,
					BackgroundImage: user[i].Avatar,
					Signature:       user[i].Avatar,
					WorkCount:       user[i].WorkCount,
					TotalFavorited:  user[i].TotalFavorited,
					FavoriteCount:   user[i].FavoriteCount,
					FollowerCount:   user[i].FollowerCount,
					FollowCount:     user[i].FollowCount,
					IsFollow:        isFollow[i],
				})
		}
	}

	return fullUser, true
}

func QueryUserListByMix(userId int64, searchId []int64, isUserListCache []bool) ([]*api.FullUser, bool) {
	fmt.Println("find list user_id mix: ", userId)

	var cacheUserIdList, modelUserIdList []int64
	for i, isCache := range isUserListCache {
		if isCache {
			cacheUserIdList = append(cacheUserIdList, searchId[i])
		} else {
			modelUserIdList = append(modelUserIdList, searchId[i])
		}
	}

	//这里做异步并发
	//查redis
	cacheUserList, ok := QueryUserListByCache(userId, cacheUserIdList)
	//查model层
	modelUserList, _ := QueryUserListByModel(userId, modelUserIdList)

	if !ok {
		return make([]*api.FullUser, len(searchId)), false
	}

	cacheUserMap := make(map[int64]*api.FullUser)
	modelUserMap := make(map[int64]*api.FullUser)
	for _, value := range cacheUserList {
		if value == nil {
			continue
		}
		cacheUserMap[value.Id] = value
	}
	for _, value := range modelUserList {
		if value == nil {
			continue
		}
		modelUserMap[value.Id] = value
	}

	var userList []*api.FullUser
	userList = make([]*api.FullUser, len(searchId))

	for i, id := range searchId {
		if _, ok = cacheUserMap[id]; ok {
			//来源于缓存
			userList[i] = cacheUserMap[id]
		} else if _, ok = modelUserMap[id]; ok {
			//来源于Model层
			userList[i] = modelUserMap[id]
		} else {
			userList[i] = nil
		}
	}
	return userList, true
}
