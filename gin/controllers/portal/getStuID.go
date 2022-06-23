package portal

import (
	"NCShiftSystem/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetStuID 根据 Unique 键，返回用户学号
func (con Controller) GetStuID(c *gin.Context) {
	val := c.Query("value")
	key := c.Query("key")
	fmt.Println("val:", val, "key:", key)
	stuID := tool.GetStuID(key, val)
	if stuID != ""{
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"stuID": stuID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	}

}