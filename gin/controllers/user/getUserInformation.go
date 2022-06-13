package user

import (
	"NCShiftSystem/jwt"
	"NCShiftSystem/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Controller struct {
}

// GetUserInformation 解析token获取用户信息返回
func (con Controller) GetUserInformation(c *gin.Context) {
	authHeader := c.Request.Header.Get("Cookie")
	parts := strings.Split(authHeader, "token=")
	token := parts[1]
	claims, err := jwt.ParseToken(token)
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"err": err,
		})
	} else {
		user := tool.GetUserInfo(claims.StuID)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"stuID": user.Stuid,
			"name": user.Name,
			"email": user.Email,
			"nickname": user.Nickname,
			"telephone": user.Telephone,
		})
	}
}