package utils

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 定义过期时间
const TokenExpireDuration = time.Hour * 24

var MySecret = []byte("这是一段生成token的密钥")

// 用来决定JWT中应该存储哪些数据，username是自定义数据
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//生成token并返回
func GenToken(username string) (string, error) {
	c := MyClaims{
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
		//token存储在Authorization中
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2,
				"msg":  "请求头中的token为空",
			})
			//终止后续的请求
			c.Abort()
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "请求头中的auth格式错误",
			})
			c.Abort()
			return
		}
		//parts[1]存储着用户的信息--用户名
		info, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "无效的token",
			})
			c.Abort()
			return
		}

		//保存当前请求信息到上下文c中
		c.Set("username", info.Username)
		//继续执行后续的请求
		c.Next()

	}
}
