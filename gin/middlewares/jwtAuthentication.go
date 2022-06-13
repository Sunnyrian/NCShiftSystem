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
		authHeader := c.Request.Header.Get("Cookie")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg": "无权限访问，请先登录",
			})
			c.Abort()
			return
		}


		// 按空格拆分
		token := strings.Split(authHeader, "token=")
		log.Print("token:", token[1])
		if token[1] == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg": "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		//解析 token 包含的信息
		claims, err := jwt.ParseToken(token[1])
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
