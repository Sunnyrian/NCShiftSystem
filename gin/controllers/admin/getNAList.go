package admin

import (
	"NCShiftSystem/model"
	"github.com/gin-gonic/gin"
	"net/http"
)


// GetNAList 向前端返回 json 格式网管列表
func (con Controller) GetNAList(c *gin.Context) {
	db := model.DB
	var user []model.User
	db.Select("stuid", "nickname", "name", "email", "telephone", "status").Find(&user)
	//result := db.Find(&user)
	//for i := int64(0); i < result.RowsAffected ; i ++ {
	//	fmt.Println(user[i].ID)
	//}

	// 这里不对 user 再进一步处理， 会有损耗， password = ''
	c.JSON(http.StatusOK, user)
}