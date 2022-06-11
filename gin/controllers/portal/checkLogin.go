package portal

import (
	"NCShiftSystem/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type checkLoginForm struct {
	Token string `json:"token"`
}

func (con Controller) CheckLogin(c *gin.Context) {
	var form checkLoginForm
	err := c.ShouldBind(&form)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(form.Token)
		_, err := jwt.ParseToken(form.Token)
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