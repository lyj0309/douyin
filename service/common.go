package service

import (
	"github.com/lyj0309/douyin/db"
	"strconv"
)

type UserRes struct {
	*subUser
	IsFollow bool `json:"is_follow,omitempty"`
}

type subUser struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
}

// GetUser 获取数据库里面的user，返回user结构体
func GetUser(userID string) *db.User {
	var user db.User
	id, _ := strconv.Atoi(userID)
	db.Mysql.Where(`id = ?`, id).First(&user)
	return &user
}

//GetUserRes 获取返回json的user结构体
func GetUserRes(MyUserID string, toUserID []uint) *[]UserRes {
	var subUsers []subUser
	var res []UserRes

	//fmt.Println(MyUserID, toUserID)
	db.Mysql.Model(&db.User{}).Find(&subUsers, toUserID)
	//fmt.Println(subUsers)
	for i := 0; i < len(subUsers); i++ {
		res = append(res, UserRes{
			subUser:  &subUsers[i],
			IsFollow: IsFollow(MyUserID, Itoa(subUsers[i].Id)),
		})
	}

	return &res
}

// IsFollow 是否关注，传入两个userID,一个是自己的，一个是是否关注的
func IsFollow(uid, toUid string) bool {
	k := followPrefix + uid

	return db.Rdb.SIsMember(ctx, k, toUid).Val()
}

//Itoa 将int类型转成string
func Itoa[V int | uint | int64](i V) string {
	return strconv.Itoa(int(i))
}
