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

	val := c.Query("val")
	key := c.Query("key")
	fmt.Println("val:", val, "key:", key)
	flag := tool.CheckUserExist(key, val)

	c.JSON(http.StatusOK, gin.H{
		"exist": flag,
	})

}
