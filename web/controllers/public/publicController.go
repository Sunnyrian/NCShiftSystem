package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PublicController struct {

}

func (con PublicController) Success(c *gin.Context) {
	c.String(http.StatusOK, "成功")
}

func (con PublicController) Error(c *gin.Context) {
	c.String(http.StatusOK, "失败")
}