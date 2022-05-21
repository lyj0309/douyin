package service

import (
	"context"
	"fmt"
	"github.com/lyj0309/douyin/db"
	"gorm.io/gorm"
	"strconv"
)

var ctx = context.Background()

const (
	followPrefix   = "f"
	followerPrefix = "fer"
)

// RelationAction 关注
func RelationAction(userID, toUserID, actionType string) error {

	uid, _ := strconv.Atoi(userID)
	tuid, _ := strconv.Atoi(toUserID)
	//新增
	if actionType == "1" {
		db.Rdb.SAdd(ctx, followPrefix+userID, toUserID)
		db.Rdb.SAdd(ctx, followerPrefix+toUserID, userID)

		db.Mysql.Model(db.User{}).Where(`id = ?`, uid).Update(`follow_count`, gorm.Expr("follow_count+1"))
		db.Mysql.Model(db.User{}).Where(`id = ?`, tuid).Update(`follower_count`, gorm.Expr("follower_count+1"))
		// 取关
	} else {
		db.Rdb.SRem(ctx, followPrefix+userID, toUserID)
		db.Rdb.SRem(ctx, followerPrefix+toUserID, userID)
		db.Mysql.Model(db.User{}).Where(`id = ?`, uid).Update(`follow_count`, gorm.Expr("follow_count-1"))
		db.Mysql.Model(db.User{}).Where(`id = ?`, tuid).Update(`follower_count`, gorm.Expr("follower_count-1"))

	}
	return nil
}

type TempUser struct {
	Name string
	ID   uint
}

func FollowList(userID string) *[]UserRes {
	return fList(userID, false)
}

func FollowerList(userID string) *[]UserRes {
	return fList(userID, false)
}

func fList(userID string, follower bool) *[]UserRes {

	var key string
	if follower {
		key = followerPrefix + userID
	} else {
		key = followPrefix + userID
	}
	uidStrs := db.Rdb.SMembers(ctx, key).Val()
	fmt.Println(uidStrs)

	var uids []uint
	for _, str := range uidStrs {
		i, _ := strconv.Atoi(str)
		uids = append(uids, uint(i))
	}

	return GetUserRes(userID, uids)
}
