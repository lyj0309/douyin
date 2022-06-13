package db

import (
	"strconv"
	"testing"
	"time"
)

func TestDB(t *testing.T) {
	for i := 1; i <= 10; i++ {
		Mysql.Create(&User{
			ID:           uint(i),
			Name:         strconv.Itoa(i) + "aa",
			Password:     strconv.Itoa(i) + "aa",
			RegisterTime: time.Now(),
		})
	}
}
