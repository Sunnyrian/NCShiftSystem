package portal

import (
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
	fmt.Println("233")
	fmt.Println(form)
	if err != nil {
		fmt.Println(err)
	} else {
		db.Where("stuid = ?", form.StuID).Find(&user)
		if user.Password == form.Password {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
			})
		}
	}
}