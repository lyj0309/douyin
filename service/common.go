package service

import (
	"github.com/lyj0309/douyin/db"
	"strconv"
)

func GetUser(userID string) *db.User {
	var user db.User
	id, _ := strconv.Atoi(userID)
	db.Mysql.Where(`id = ?`, id).First(&user)
	return &user
}
