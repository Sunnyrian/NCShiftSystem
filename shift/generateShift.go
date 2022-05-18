package shift

import (
	shiftConst "NCShiftSystem/const"
	"NCShiftSystem/model"
	"NCShiftSystem/tool"
	"fmt"
	"gorm.io/datatypes"
	"strconv"
)

/*
   -1 不值班
   0 10:00 - 12:00 14:00-16:00 16:00-18:00 18:00-21:00 4个时间段 0 1 2 3
   1 10:00 - 12:00 18:00-21:00 						   2个时间段 0 3
*/

//根据前端页面传回的值，生成对应的值班表
func GenerateShift(statDate string, endDate string, numPerShift string, week [7]string) {
	startWeek, startWeekDay := tool.DateToWeekDay(statDate)
	endWeek, endWeekDay := tool.DateToWeekDay(endDate)
	numberPerShift, err := strconv.Atoi(numPerShift)
	var shift []model.Shift
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < numberPerShift; i++ {
		// j 是第几周 k 是星期
		//fmt.Println("startWeek:",startWeek,"endWeek:",endWeek)
		for j := startWeek; j <= endWeek; j++ {
			// init
			var k = -1
			var dayOfWeek = -1
			if j == startWeek {
				k = startWeekDay
			} else {
				k = 0
			}
			if j == endWeek {
				dayOfWeek = endWeekDay
			} else {
				dayOfWeek = shiftConst.DayOfWeek
			}
			for ; k < dayOfWeek || k == 0 && dayOfWeek == 0; k++ {
				if week[k] == "-1" {

				} else if week[k] == "0" {
					for m := 0; m < 4; m++ {
						//fmt.Println("i:",i,"j:",j,"k:",k)
						d := tool.WeekDayToDate(j, k)
						singleShift := model.Shift{UserID: 0, Week: j, Weekday: k, TimePeriod: m, Date: datatypes.Date(d), Status: -1}
						shift = append(shift, singleShift)
					}
				} else if week[k] == "1" {
					//fmt.Println("i:",i,"j:",j,"k:",k)
					d := tool.WeekDayToDate(j, k)
					singleShift := model.Shift{UserID: 0, Week: j, Weekday: k, TimePeriod: 0, Date: datatypes.Date(d), Status: -1}
					singleShift2 := model.Shift{UserID: 0, Week: j, Weekday: k, TimePeriod: 3, Date: datatypes.Date(d), Status: -1}
					shift = append(shift, singleShift)
					shift = append(shift, singleShift2)
				}
			}
		}
	}
	//写入数据库
	//fmt.Println(shift)
	db := model.ConnectToMySQL()
	db.Create(&shift)
}
