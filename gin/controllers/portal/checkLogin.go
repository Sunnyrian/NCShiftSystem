package portal

import (
	"NCShiftSystem/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func (con Controller) CheckLogin(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"success:": false,
			"err": err,
		})
	} else {
		_, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"err": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		}
	}
}