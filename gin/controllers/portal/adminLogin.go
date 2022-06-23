package portal

import (
	"NCShiftSystem/jwt"
	"NCShiftSystem/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type adminLoginForm struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func (con Controller) AdminLogin(c *gin.Context) {
	db := model.DB
	var form adminLoginForm
	err := c.ShouldBind(&form)
	var admin model.AdminUser
	if err != nil {
		fmt.Println(err)
	} else {
		db.Where("nickname = ?", form.Nickname).Find(&admin)
		if admin.Password == form.Password {
			token, err := jwt.GenerateToken(form.Nickname)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg": err,
				})
				return
			}
			c.SetCookie("token", token, 10800, "", "", false, false)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"token": token,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
			})
		}
	}

}