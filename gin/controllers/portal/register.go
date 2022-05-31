package portal

import (
	shiftConst "NCShiftSystem/const"
	"NCShiftSystem/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (con Controller) Register(c *gin.Context) {
	db := model.DB
	nickname := c.Query("nickname")
	password := c.Query("password")
	name := c.Query("name")
	stuID := c.Query("stuID")
	telephone := c.Query("telephone")
	email := c.Query("email")
	status := shiftConst.DefaultStatus
	fmt.Println("nickname:"+nickname)
	user := model.User{
		Name: name,
		Nickname: nickname,
		Password: password,
		Stuid: stuID,
		Telephone: telephone,
		Email: email,
		Status: status,
	}
	fmt.Println(user)
	db.Create(&user)
}