package portal

import (
	shiftConst "NCShiftSystem/const"
	"NCShiftSystem/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


type registerForm struct {
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Password string `json:"password"`
	StuID string `json:"stuID"`
	Telephone string `json:"telephone"`
	Email string `json:"email"`
}

func (con Controller) Register(c *gin.Context) {
	db := model.DB
	var form registerForm
	err := c.ShouldBind(&form)
	if err != nil || form.Name == ""{
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
		})
	} else {
		status := shiftConst.DefaultStatus
		user := model.User{
			Name: form.Name,
			Nickname: form.Nickname,
			Password: form.Password,
			Stuid: form.StuID,
			Telephone: form.Telephone,
			Email: form.Email,
			Status: status,
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		fmt.Println(user)
		db.Create(&user)
	}
}
