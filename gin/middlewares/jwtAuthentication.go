package middlewares

import (
	"NCShiftSystem/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg": "无权限访问，请先登录",
			})
			c.Abort()
			return
		}
		log.Print("token", authHeader)

		// 按空格拆分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg": "请求头中auth格式有误",
			})
			c.Abort()
			return
		}

		// 解析 token 包含的信息
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg": "无效的 Token",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()

	}
}
