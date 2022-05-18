package main

import (
	"NCShiftSystem/model"
	"NCShiftSystem/shift"
	"fmt"
	"time"
)

/* 数据库，Go 程序中，所有数都从 0 开始 */
func main() {
	var startTime = time.Now()

	db := model.ConnectToMySQL()
	//根据数据库中的网管列表，在本地目录中读取 Xls，更新网管的 OT 表
	//shift.InitOccupation(db)
	//readXls.ReadAllXls(db)

	//生成排班表
	//week := [7]string{"0", "0", "0", "0", "0", "1", "1"}
	//shift.GenerateShift("2022-05-16", "2022-06-20", "2", week)

	//排班
	startDate := "2022-05-16"
	endDate := "2022-06-20"
	shift.Shift(db, startDate, endDate)

	//web.StartWeb()
	var end = time.Now()
	fmt.Println("程序运行了:", end.Sub(startTime))
}

// 测试 DateToWeekDay
//week, weekday := tool.DateToWeekDay("2022-05-22")
//fmt.Println(week, weekday)

// 测试 WeekDayToDate
//d := tool.WeekDayToDate(12, 0)
//fmt.Println(d)

// 测试 GenerateShift
//week := [7]string{"0", "0", "0", "0", "0", "1", "1"}
//shift.GenerateShift("2022-05-16", "2022-05-22", "2", week)

// 测试 StringDateToGromDate
//s := tool.StringDateToGromDate("2022-05-16")
//fmt.Println(s)
