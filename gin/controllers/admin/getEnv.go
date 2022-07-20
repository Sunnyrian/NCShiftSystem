package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

type dotEnv struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

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
	var jsonEnv []dotEnv
	for k,v := range Env{
		oneEnv := dotEnv{
			Key: k,
			Value: v,
		}
		jsonEnv = append(jsonEnv, oneEnv)
	}

	c.JSON(http.StatusOK, jsonEnv)
}