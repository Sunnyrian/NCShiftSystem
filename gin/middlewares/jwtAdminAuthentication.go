package middlewares

import (
	"NCShiftSystem/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Cookie")

		token := strings.Split(authHeader, "token=")

		claims, err := jwt.ParseToken(token[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg": err,
			})
		}

		if !claims.Admin {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg": "当前登录账户非管理员",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}