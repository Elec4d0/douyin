package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"strconv"
	"time"
	"userInfo/services/protos/kitex_gen/api"
)

func InsertNilUserRedis(id int64) error {
	err := redisClient.HSet(ctx, strconv.FormatInt(id, 10),
		"Id", id,
	).Err()
	rand.NewSource(time.Now().UnixNano())
	r := rand.Intn(30)
	_ = redisClient.Expire(ctx, strconv.FormatInt(id, 10), 300*time.Second+time.Duration(r)*time.Second)
	log.Println("存入空值！")
	if err != nil {
		fmt.Println("缓存存储错误！")
		fmt.Println(err)
		return err
	}
	fmt.Println("缓存存储成功！")
	return nil
}

func InsertUserRedis(user *api.FullUser) error {
	err := redisClient.HSet(ctx, strconv.FormatInt(user.Id, 10),
		"Id", user.Id,
		"Name", user.Name,
		"FollowCount", *user.FollowCount,
		"FollowerCount", *user.FollowerCount,
		"Avatar", *user.Avatar,
		"BackgroundImage", *user.BackgroundImage,
		"Signature", *user.Signature,
		"TotalFavorited", *user.TotalFavorited,
		"WorkCount", *user.WorkCount,
		"FavoriteCount", *user.FavoriteCount,
	).Err()
	rand.NewSource(time.Now().UnixNano())
	r := rand.Intn(30)
	_ = redisClient.Expire(ctx, strconv.FormatInt(user.Id, 10), 300*time.Second+time.Duration(r)*time.Second)
	log.Println("待缓存数据：", *user.FavoriteCount)
	if err != nil {
		fmt.Println("缓存存储错误！")
		fmt.Println(err)
		return err
	}
	fmt.Println("缓存存储成功！")
	return nil
}

func InsertUserListRedis(userList []*api.FullUser) error {
	pipeline := redisClient.Pipeline()
	for _, value := range userList {
		if value == nil || value.Id == 0 {
			continue
		}
		pipeline.HSet(ctx, strconv.FormatInt(value.Id, 10),
			"Id", value.Id,
			"Name", value.Name,
			"FollowCount", *value.FollowCount,
			"FollowerCount", *value.FollowerCount,
			"Avatar", *value.Avatar,
			"BackgroundImage", *value.BackgroundImage,
			"Signature", *value.Signature,
			"TotalFavorited", *value.TotalFavorited,
			"WorkCount", *value.WorkCount,
			"FavoriteCount", *value.FavoriteCount,
		)
		rand.NewSource(time.Now().UnixNano())
		r := rand.Intn(30)
		pipeline.Expire(ctx, strconv.FormatInt(value.Id, 10), 300*time.Second+time.Duration(r)*time.Second)
	}

	cmders, err := pipeline.Exec(ctx)
	if err != nil {
		log.Println("pipeline exec failed. err ", err)
		return err
	}
	cnt := 0
	for _, cmder := range cmders {
		if cnt%2 == 1 {
			continue
		}
		cnt += 1
		cmd := cmder.(*redis.IntCmd)
		_, err = cmd.Result()
		if err != nil {
			log.Println("err ", err)
		}
	}
	fmt.Println("批量缓存存储成功！")
	return nil
}

func FindUserRedis(id int64) (*User, error) {
	// 外部判断缓存是否存在
	mp, err := redisClient.HGetAll(ctx, strconv.FormatInt(id, 10)).Result()
	if err != nil {
		fmt.Println("缓存查询错误！")
		fmt.Println(err)
		return nil, err
	}
	if len(mp) <= 1 {
		fmt.Println("缓存值空！")
		fmt.Println(err)
		return nil, err
	}
	followCount := str2int64(mp["FollowCount"])
	followerCount := str2int64(mp["FollowerCount"])
	avatar := mp["Avatar"]
	backgroundImage := mp["BackgroundImage"]
	signature := mp["Signature"]
	totalFavorited := str2int64(mp["TotalFavorited"])
	workCount := str2int64(mp["WorkCount"])
	favoriteCount := str2int64(mp["FavoriteCount"])
	user := &User{
		Id:              str2int64(mp["Id"]),
		Name:            mp["Name"],
		FollowCount:     &followCount,
		FollowerCount:   &followerCount,
		Avatar:          &avatar,
		BackgroundImage: &backgroundImage,
		Signature:       &signature,
		TotalFavorited:  &totalFavorited,
		WorkCount:       &workCount,
		FavoriteCount:   &favoriteCount,
	}
	fmt.Println("缓存查询成功！")
	log.Println(user)
	return user, nil
}

func FindUserListRedis(id []int64) ([]*User, error) {
	pipeline := redisClient.Pipeline()
	var userList []*User
	for _, value := range id {
		pipeline.HGetAll(ctx, strconv.FormatInt(value, 10))
	}

	cmders, err := pipeline.Exec(ctx)
	if err != nil {
		log.Println("pipeline exec failed. err ", err)
		return make([]*User, len(id)), err
	}
	for _, cmder := range cmders {
		cmd := cmder.(*redis.MapStringStringCmd)
		data, err := cmd.Result()
		if err != nil {
			log.Println("err ", err)
		}
		if len(data) <= 1 {
			userList = append(userList, nil)
		} else {
			followCount := str2int64(data["FollowCount"])
			followerCount := str2int64(data["FollowerCount"])
			avatar := data["Avatar"]
			backgroundImage := data["BackgroundImage"]
			signature := data["Signature"]
			totalFavorited := str2int64(data["TotalFavorited"])
			workCount := str2int64(data["WorkCount"])
			favoriteCount := str2int64(data["FavoriteCount"])
			userList = append(userList,
				&User{
					Id:              str2int64(data["Id"]),
					Name:            data["Name"],
					FollowCount:     &followCount,
					FollowerCount:   &followerCount,
					Avatar:          &avatar,
					BackgroundImage: &backgroundImage,
					Signature:       &signature,
					TotalFavorited:  &totalFavorited,
					WorkCount:       &workCount,
					FavoriteCount:   &favoriteCount,
				})
		}

	}
	fmt.Println("批量缓存查询成功！")
	return userList, nil
}

func CheckUserExist(id int64) bool {
	isUserExist, err := redisClient.HExists(ctx, strconv.FormatInt(id, 10), "Id").Result()
	reply := redisClient.TTL(ctx, strconv.FormatInt(id, 10))
	var time = 10 * time.Second
	log.Println("time:  ", reply.Val())
	if err != nil || reply.Val() < time {
		return false
	}
	return isUserExist
}

func CheckUserListExist(id []int64) ([]bool, int64) {
	var isUserListExist []bool
	// 0: model, 1: mix, 2: cache
	var flag int64
	var trueCnt int64
	var falseCnt int64
	pipeline := redisClient.Pipeline()
	for _, value := range id {
		pipeline.HExists(ctx, strconv.FormatInt(value, 10), "Id").Result()
		pipeline.TTL(ctx, strconv.FormatInt(value, 10))
	}

	cmders, err := pipeline.Exec(ctx)
	if err != nil {
		log.Println("pipeline exec failed. err ", err)
		return make([]bool, len(id)), 0
	}

	for index, cmder := range cmders {
		if index%2 == 1 {
			continue
		}
		cmd := cmder.(*redis.BoolCmd)
		data, err := cmd.Result()
		if err != nil {
			falseCnt += 1
			isUserListExist = append(isUserListExist, false)
			log.Println("err ", err)
			falseCnt += 1
			continue
		}

		ttl := cmders[index+1].(*redis.DurationCmd).Val()
		log.Println("time: ", ttl)
		if ttl < time.Second*10 {
			data = false
		}
		isUserListExist = append(isUserListExist, data)
		if data == true {
			trueCnt += 1
		} else {
			falseCnt += 1
		}
	}
	if falseCnt != 0 && trueCnt != 0 {
		flag = 1
	} else if falseCnt == 0 && trueCnt != 0 {
		flag = 2
	} else {
		flag = 0
	}
	return isUserListExist, flag
}

func str2int64(str string) int64 {
	integer, err := strconv.Atoi(str)
	if err != nil {
		return int64(0)
	}
	return int64(integer)
}
