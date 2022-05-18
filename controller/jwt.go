package controller

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var AppSecret = "fj4i448g5sdffa8wef6awe4f68" //viper.GetString会设置这个值(32byte长度)
var AppIss = "fake_douyin"                   //这个值会被viper.GetString重写

func JwtGenerateToken(userID uint, d time.Duration) (string, error) {
	expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        strconv.Itoa(int(userID)),
		Issuer:    AppIss,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, stdClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(AppSecret))
	if err != nil {
		logrus.WithError(err).Fatal("config is wrong, can not generate jwt")
	}
	return tokenString, err
}

//JwtParseUser 解析payload的内容,得到用户信息
//gin-middleware 会使用这个方法
func JwtParseUser(tokenString string) (string, error) {
	if tokenString == "" {
		return "", errors.New("no token is found in Authorization Bearer")
	}
	claims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(AppSecret), nil
	})
	if err != nil {
		return "", err
	}
	return claims.Id, err

}
