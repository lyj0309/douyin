package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const ctxUidKey = "user_id"

func Auth(c *gin.Context) {
	token, ok := c.GetQuery("token")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "token未找到"})
		return
	}
	uid, err := JwtParseUser(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	//store the user Model in the context
	c.Set(ctxUidKey, uid)
	c.Next()
}
