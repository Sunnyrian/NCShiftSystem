package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

// GetEnv 向前端返回当前的环境变量
func (con Controller) GetEnv(c *gin.Context){
	Env, err := godotenv.Read()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	}
	// 将 map 转换为 json 格式
	bytes,_ := json.Marshal(Env)
	jsonEnv := string(bytes)
	c.JSON(http.StatusOK, jsonEnv)
}