package admin

import (
	"NCShiftSystem/const"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type envForm struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

// ModifyEnv 修改后端的环境变量
func (con Controller) ModifyEnv(c *gin.Context){
	var form envForm
	err := c.ShouldBind(&form)
	if err != nil {
		log.Println(err)
	}
	if shiftConst.ModifyEnv(form.Key, form.Value) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	}
}