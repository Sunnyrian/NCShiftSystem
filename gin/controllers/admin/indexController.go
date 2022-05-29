package admin

import (
	shiftConst "NCShiftSystem/const"
	"net/http"

	//"NCShiftSystem/web/controllers/public"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IndexController struct{
	//public.PublicController
}

//// Index 返回主页面
//func (con IndexController) Index(c *gin.Context) {
//	//con.Success(c)
//	c.HTML(http.StatusOK, "admin/index.html",gin.H{})
//}

// Shift Post方法传到后端排班参数
func (con IndexController) Shift(c *gin.Context) {
	// fixbug 一定要定义好数组长度
	//weekday or holiday or not need to work
	var week [shiftConst.DayOfWeek]string
	for i := 0; i < shiftConst.DayOfWeek; i++ {
		s := strconv.Itoa(i)
		week[i] = c.PostForm(s)
	}
	startDate := c.PostForm("startDate")
	endDate := c.PostForm("endDate")
	numberPerShift := c.PostForm("numPerShift")
	//东西都拿到了，接下来开始排班吧
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"startDate" : startDate,
		"endDate" : endDate,
		"0" : week[0],
		"1" : week[1],
		"2" : week[2],
		"3" : week[3],
		"4" : week[4],
		"5" : week[5],
		"6" : week[6],
		"numberPerShift" : numberPerShift,
	})
	//shift.GenerateShift(startDate, endDate, numberPerShift, week)
}