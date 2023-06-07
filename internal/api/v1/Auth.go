package v1

import (
	"library-sys-go/pkg/resp"

	"github.com/gin-gonic/gin"
)

var auth = "STFNg5PqqXXgUAd5hTzjj3qrJzhunopXUjUnF7C4"

func Auth() gin.HandlerFunc {
	// 检查Header Authorization
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			resp.Error(c, 401, "Un Authorization")
			c.Abort()
			return
		}
		// 检查token是否有效
		if token != auth {
			resp.Error(c, 401, "token is invalid")
			c.Abort()
			return
		}
		c.Next()
	}
}

func Ping(c *gin.Context) {
	resp.SuccessData(c, "pong")
}
