package portal

import (
	"NCShiftSystem/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {

}

// CheckUserExist 查询数据库该用户有没有存在
func (con Controller) CheckUserExist(c *gin.Context) {

	val := c.PostForm("value")
	key := c.PostForm("key")
	//val := c.GetString("val")
	//key := c.GetString("key")
	fmt.Println(c)
	fmt.Println(c.GetRawData())
	fmt.Println("val:",val,"key:",key)
	flag := tool.CheckUserExist(val, key)

	c.JSON(http.StatusOK, gin.H{
		"val": val,
		"key": key,
		"exist": flag,
	})
}