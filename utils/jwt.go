package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 定义过期时间
const TokenExpireDuration = time.Hour * 24

var MySecret = []byte("这是一段生成token的密钥")

// 用来决定JWT中应该存储哪些数据
type MyClaims struct {
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

//生成token并返回
func GenToken(userId uint, username string) (string, error) {
	c := MyClaims{
		userId,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "userFunction",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

//解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, err
	}
	//token失效
	return nil, errors.New("invalid token")
}

// 后续会携带着token进行请求接口

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		token, ok := c.GetQuery("token")

		if !ok {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  "未携带token",
			})
			return
		}

		res, err := ParseToken(token)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  err.Error(),
			})
			return
		}
		//保存当前请求信息到上下文c中
		c.Set("username", res.Username)
		c.Set("user_id", res.UserId)
		//继续执行后续的请求
		c.Next()

	}
}
