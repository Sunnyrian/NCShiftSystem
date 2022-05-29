package portal

import "github.com/gin-gonic/gin"

// checkUserExist 查询数据库该用户有没有存在
func checkUserExist(c *gin.Context) {
	nickname := c.PostForm("nickname")

}