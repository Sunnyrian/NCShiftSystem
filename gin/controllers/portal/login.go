package portal

import (
	"NCShiftSystem/jwt"
	"NCShiftSystem/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginForm struct {
	StuID string `json:"stuID"`
	Password string `json:"password"`
}

func (con Controller) Login(c *gin.Context) {
	db := model.DB
	var form loginForm
	err := c.ShouldBind(&form)
	var user model.User
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	} else {
		db.Where("stuid = ?", form.StuID).Find(&user)
		if user.Password == form.Password {
			token, err := jwt.GenerateToken(form.StuID, false)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"success": false,
					"msg": err,
				})
				return
			}
			c.SetCookie("token", token, 10800, "", "", false, false)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				//"token": token,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
			})
		}
	}

}