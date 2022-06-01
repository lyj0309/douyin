package service

import (
	"context"
	"errors"
	"github.com/lyj0309/douyin/db"
	"github.com/sirupsen/logrus"
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

	if uid == 0 || tuid == 0 {
		return errors.New("userid为0")
	}

	//开始事务
	err := db.Mysql.Transaction(func(tx *gorm.DB) error {
		var err error
		//新增
		if actionType == "1" {
			db.Rdb.SAdd(ctx, followPrefix+userID, toUserID)
			db.Rdb.SAdd(ctx, followerPrefix+toUserID, userID)

			err = tx.Model(db.User{}).Where(`id = ?`, uid).Update(`follow_count`, gorm.Expr("follow_count+1")).Error
			if err != nil {
				return err
			}
			err = tx.Model(db.User{}).Where(`id = ?`, tuid).Update(`follower_count`, gorm.Expr("follower_count+1")).Error
			if err != nil {
				return err
			}
			// 取关
		} else {
			db.Rdb.SRem(ctx, followPrefix+userID, toUserID)
			db.Rdb.SRem(ctx, followerPrefix+toUserID, userID)
			err = tx.Model(db.User{}).Where(`id = ?`, uid).Update(`follow_count`, gorm.Expr("follow_count-1")).Error
			if err != nil {
				return err
			}

			err = tx.Model(db.User{}).Where(`id = ?`, tuid).Update(`follower_count`, gorm.Expr("follower_count-1")).Error
			if err != nil {
				return err
			}

		}
		logrus.Info("用户关注：事务提交", tuid, uid)

		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		return err
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
	return fList(userID, true)
}

func fList(userID string, follower bool) *[]UserRes {

	var key string
	if follower {
		key = followerPrefix + userID
	} else {
		key = followPrefix + userID
	}
	uidStrs := db.Rdb.SMembers(ctx, key).Val()
	//fmt.Println(uidStrs, key, userID, follower)

	var uids []uint
	for _, str := range uidStrs {
		i, _ := strconv.Atoi(str)
		uids = append(uids, uint(i))
	}

	return GetUserRes(userID, uids)
}
